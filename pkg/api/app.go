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

package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/kube-all/mock-runner/pkg/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"k8s.io/klog/v2"
	"net/http"
)

func app() {
	prometheus.Init()
	ws := new(restful.WebService)
	container := restful.NewContainer()
	ws.Route(ws.GET("/").To(index))
	ws.Route(ws.GET("/metrics").To(metrics))
	container.Add(ws)
	klog.V(1).Infof("mock app server will start with port 8081")
	server := &http.Server{Addr: ":8081", Handler: container}
	klog.Fatal("app server", server.ListenAndServe())
}
func index(req *restful.Request, resp *restful.Response) {
	resp.Write([]byte( "this is mock app server"))
}
func metrics(req *restful.Request, resp *restful.Response) {
	promhttp.Handler().ServeHTTP(resp.ResponseWriter, req.Request)
}
