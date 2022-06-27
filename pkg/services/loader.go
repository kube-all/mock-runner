/*
Copyright 2022 The kubeall.com Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package services

import (
	"encoding/json"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/kube-all/mock-runner/cmd/server/options"
	"github.com/kube-all/mock-runner/pkg/core"
	"github.com/kube-all/mock-runner/pkg/prometheus"
	"github.com/kube-all/mock-runner/pkg/utils"
	"gopkg.in/yaml.v3"
	"k8s.io/klog/v2"
	"net/http"
	"path"
	"strings"
)

type MockServer struct {
	Option    *options.Options
	Container *restful.Container
}

func (m *MockServer) LoadAPI() {
	apiMap := utils.LoadDirFileData(path.Join(m.Option.Path, "apis"))
	ws := new(restful.WebService)
	var (
		err error
	)
	for apiPath, data := range apiMap {
		prometheus.ApiTotal.Add(1)
		var api core.APIDefinition
		if strings.HasSuffix(apiPath, ".json") {
			err = json.Unmarshal(data, &api)
			if err != nil {
				klog.Errorf("json unmarshal: %s failed, err: %s", apiPath, err.Error())
				continue
			}
		} else if strings.HasSuffix(apiPath, ".yaml") || strings.HasSuffix(apiPath, ".yml") {
			err = yaml.Unmarshal(data, &api)
			if err != nil {
				klog.Errorf("yaml unmarshal: %s failed, err: %s", apiPath, err.Error())
				continue
			}
		}
		api.DefaultValue()
		if errs := api.Validator(); len(errs) > 0 {
			klog.Errorf("api: [%s] validate failed, errs: [%s]", api.Spec.Name, strings.Join(errs, ";"))
			continue
		}
		api.Spec.Method = strings.ToUpper(api.Spec.Method)
		// format api definition
		var rb *restful.RouteBuilder
		switch api.Spec.Method {
		case http.MethodGet:
			rb = ws.GET(api.Spec.Path)
		case http.MethodPost:
			rb = ws.POST(api.Spec.Path)
		case http.MethodPut:
			rb = ws.PUT(api.Spec.Path)
		case http.MethodDelete:
			rb = ws.DELETE(api.Spec.Path)
		case http.MethodPatch:
			rb = ws.PATCH(api.Spec.Path)
		default:
			klog.Errorf("not support method: %s for api: %s, path; %s", api.Spec.Method, api.Spec.Name, api.Spec.Path)
			continue
		}
		rb.Consumes(api.Spec.Consumes...).Produces(api.Spec.Produces...)
		rb.Doc(api.Spec.Description).Metadata(restfulspec.KeyOpenAPITags, api.Spec.Tags)
		for _, param := range api.Spec.Parameters {
			switch param.Position {
			case "query":
				rb.Param(ws.QueryParameter(param.Name, param.Description).DataType(param.Type).Required(param.Required))
			case "path":
				rb.Param(ws.PathParameter(param.Name, param.Description).DataType(param.Type).Required(true))
			case "header":
				rb.Param(ws.HeaderParameter(param.Name, param.Description).DataType(param.Type).Required(param.Required))
			default:
				klog.Warningf("unknown param position: %s", param.Position)
			}
		}
		if api.Spec.Method == http.MethodPatch || api.Spec.Method == http.MethodPost || api.Spec.Method == http.MethodPut {
		}
		ws.Route(rb.To(restful.RouteFunction(APIHandler(&api))))
		// 增加计数
		prometheus.ApiSuccessTotal.Add(1)
	}
	m.Container.Add(ws)
}
