package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/go-openapi/spec"
	"github.com/kube-all/mock-runner/cmd/server/options"
	"github.com/kube-all/mock-runner/pkg/embeds"
	"github.com/kube-all/mock-runner/pkg/services"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"k8s.io/klog/v2"
	"net/http"
	"path"
)

func mock(o *options.Options) {
	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.DefaultResponseContentType(restful.MIME_JSON)
	container := restful.DefaultContainer
	container.Filter(container.OPTIONSFilter)
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{},
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		CookiesAllowed: true,
		Container:      container}
	container.Filter(cors.Filter)

	//load api
	mockSvc := services.MockServer{
		Option:    o,
		Container: container,
	}
	if data, err := ioutil.ReadFile(path.Join(o.Path, "config.yaml")); err == nil {
		yaml.Unmarshal(data, &mockSvc.Config)
	} else {
		klog.Warningf("can't get config from path: %s, err: %s", path.Join(o.Path, "config.yaml"), err.Error())
	}
	mockSvc.LoadAPI()

	//swagger
	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject(mockSvc.Config.Tags),
	}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))
	http.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(embeds.StaticFileSystem())))

	klog.V(1).Infof("mock server will start with port 8080")
	klog.Fatal(http.ListenAndServe(":8080", nil))
}

func enrichSwaggerObject(tags map[string]string) restfulspec.PostBuildSwaggerObjectFunc {
	return func(swo *spec.Swagger) {
		swo.Info = &spec.Info{
			InfoProps: spec.InfoProps{
				Title:       "mock",
				Description: "An Open Source Http Mock Server",
				License: &spec.License{
					LicenseProps: spec.LicenseProps{
						Name: "MIT",
						URL:  "http://mit.org"},
				},
				Version: "v1.0.0",
			},
		}
		swo.Tags = []spec.Tag{}
		for k, v := range tags {
			swo.Tags = append(swo.Tags, spec.Tag{
				TagProps: spec.TagProps{
					Name:        k,
					Description: v,
				},
			})
		}
	}
}
