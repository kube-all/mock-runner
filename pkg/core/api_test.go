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
	"github.com/emicklei/go-restful/v3"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestSimpleConditionAPIDefinition(t *testing.T) {
	def := APIDefinition{
		TypeMeta: TypeMeta{Kind: "APIDefinition", Version: "1.0"},
		Spec: &APIDefinitionSpec{
			Name:        "api demo simple condition",
			Description: "this is APIDefinition demo simple condition",
			Consumes:    []string{restful.MIME_JSON, restful.MIME_XML},
			Produces:    []string{restful.MIME_JSON, restful.MIME_XML},
			Protocol:    "http",
			Method:      "get",
			Path:        "/api/v1/demo",
			Tags:        []string{"demo"},
			Cases: []*CaseService{
				{
					Condition: &Condition{
						Simple: &SimpleCondition{
							LogicAnd: true,
							Items: []*AssertItem{
								{
									ValueFrom: "$request.query.key",
									Value:     "query-value",
									Operator:  "=",
								},
							},
						},
					},
					Response: &Response{
						Simple: &SimpleResponse{
							Header: map[string]string{"header1": "values"},
							Code:   http.StatusOK,
							Body:   SimpleResponseBody{Content: "mock response is return success"},
						},
					},
				},
			},
			Parameters: []APIParameterDefinition{
				{Position: "query", Name: "key", Description: "query param", Type: "string"},
			},
		},
	}

	data, err := yaml.Marshal(def)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("../embeds/templates/apis/demo-api-definition-simple-condition.yaml", data, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}
