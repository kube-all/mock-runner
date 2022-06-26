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

package core

import (
	"fmt"
	"github.com/kube-all/mock-runner/pkg/global"
	"net/http"
)

type CaseResponse struct {
	Code   int               `json:"code,omitempty" yaml:"code"`
	Header map[string]string `json:"header,omitempty" yaml:"header"`
	Body   []byte            `json:"body,omitempty" yaml:"body"`
}
type Response struct {
	Simple *SimpleResponse `json:"simple,omitempty" yaml:"simple,omitempty"`
}
type SimpleResponse struct {
	Code   int                `json:"code,omitempty" yaml:"code,omitempty"`
	Header map[string]string  `json:"header,omitempty" yaml:"header,omitempty"`
	Body   SimpleResponseBody `json:"body,omitempty" yaml:"body,omitempty"`
}
type SimpleResponseBody struct {
	Content     string `json:"content,omitempty" yaml:"content,omitempty"`
	RefDataFile string `json:"refDataFile,omitempty" yaml:"refDataFile,omitempty"`
}

func (c SimpleResponse) GetResponse() (response CaseResponse) {
	response.Header = c.Header
	response.Code = c.Code
	if len(c.Body.Content) > 0 {
		response.Body = []byte(c.Body.Content)
	} else if len(c.Body.RefDataFile) > 0 {
		//todo 需要进行格式转换
		if data, ok := global.MockData.Load(c.Body.RefDataFile); ok {
			response.Body = data.([]byte)
		} else {
			response.Code = http.StatusInternalServerError
			response.Body = []byte(fmt.Sprintf("can't get response body from: %s", c.Body.RefDataFile))
		}
	}
	return
}
