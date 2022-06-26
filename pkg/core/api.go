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

type TypeMeta struct {
	Kind    string `json:"kind,omitempty" yaml:"kind"`
	Version string `json:"version,omitempty"`
}
type APIDefinition struct {
	TypeMeta `json:",inline" yaml:",inline"`
	Spec     *APIDefinitionSpec `json:"spec,omitempty" yaml:"spec"`
}

type APIDefinitionSpec struct {
	// go-openapi OperationProps props
	Name        string   `json:"name,omitempty" yaml:"name"`
	Description string   `json:"description,omitempty" yaml:"description"`
	Consumes    []string `json:"consumes,omitempty" yaml:"consumes"`
	Produces    []string `json:"produces,omitempty" yaml:"produces"`
	//Schemes      []string                    `json:"schemes,omitempty" yaml:"schemes"`
	Tags []string `json:"tags,omitempty" yaml:"tags"`
	//ExternalDocs *spec.ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs"`
	Deprecated bool                     `json:"deprecated,omitempty" yaml:"deprecated"`
	Parameters []APIParameterDefinition `json:"parameters,omitempty" yaml:"parameters"`
	//Responses    *spec.Responses             `json:"responses,omitempty" yaml:"responses"`
	//  mock runner custom props
	Protocol string         `json:"protocol,omitempty" yaml:"protocol"`
	Method   string         `json:"method,omitempty" yaml:"method"`
	Path     string         `json:"path,omitempty" yaml:"path"`
	Cases    []*CaseService `json:"cases,omitempty" yaml:"cases"`
}
type CaseService struct {
	Condition *Condition `json:"condition,omitempty" yaml:"condition,omitempty"`
	Response  *Response  `json:"response,omitempty" yaml:"response,omitempty"`
}

type APIParameterDefinition struct {
	Position     string      `json:"position,omitempty" yaml:"position"`
	Name         string      `json:"name,omitempty" yaml:"name"`
	Description  string      `json:"description,omitempty" yaml:"description"`
	DefaultValue interface{} `json:"defaultValue,omitempty" yaml:"defaultValue"`
	Type         string      `json:"type,omitempty" yaml:"type"`
	Required     bool        `json:"required,omitempty" yaml:"required"`
}
