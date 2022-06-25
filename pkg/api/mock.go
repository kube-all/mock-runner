package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/go-openapi/spec"
	"github.com/kube-all/mock-runner/cmd/server/options"
	"github.com/kube-all/mock-runner/pkg/embeds"
	"k8s.io/klog/v2"
	"net/http"
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
	//swagger
	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject,
	}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))
	http.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(embeds.StaticFileSystem())))

	klog.V(1).Infof("mock server will start with port 8080")
	klog.Fatal(http.ListenAndServe(":8080", nil))
}

//enrichSwaggerObject
func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "mock runner",
			Description: "An Open Source Http Mock Server",
			Contact: &spec.ContactInfo{
				ContactInfoProps: spec.ContactInfoProps{
					Name:  "kubeall",
					Email: "kubeall@aliyun.com",
					URL:   ""},
			},
			License: &spec.License{
				LicenseProps: spec.LicenseProps{
					Name: "MIT",
					URL:  "http://mit.org"},
			},
			Version: "v1.0.0",
		},
	}
	swo.Tags = []spec.Tag{}

}

