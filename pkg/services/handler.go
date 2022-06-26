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
	"github.com/emicklei/go-restful/v3"
	"github.com/kube-all/mock-runner/pkg/core"
	"k8s.io/klog/v2"
	"net/http"
)

type APIDefinitionHandler func(req *restful.Request, resp *restful.Response)

func APIHandler(definition *core.APIDefinition) APIDefinitionHandler {
	klog.Infof("APIHandler name: %s, method: %s, path: %s",
		definition.Spec.Name, definition.Spec.Method, definition.Spec.Path)
	return func(req *restful.Request, resp *restful.Response) () {
		klog.Infof("APIDefinitionHandler APIDefinition: %s, Method: %s, RequestURI: %s",
			definition.Spec.Name, req.Request.Method, req.Request.RequestURI)

		resp.WriteHeader(http.StatusNotImplemented)
		resp.Write([]byte("your request not match any case"))
	}

}
