package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/go-openapi/spec"
	"github.com/kube-all/mock-runner/cmd/server/options"
	"github.com/kube-all/mock-runner/pkg/embeds"
	"github.com/kube-all/mock-runner/pkg/services"
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
	mockSvc.LoadAPI()
	//static
	staticPath := path.Join(o.Path, "statics")
	_, err := ioutil.ReadDir(staticPath)
	if err == nil {
		klog.Infof("static dir: %s exist, will add static handler", staticPath)
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticPath))))
	}
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

}
