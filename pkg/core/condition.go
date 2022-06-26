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
	"github.com/emicklei/go-restful/v3"
	"k8s.io/klog/v2"
	"strings"
)

type Condition struct {
	Simple *SimpleCondition `json:"simple,omitempty" yaml:"simple,omitempty"`
}
type SimpleCondition struct {
	LogicAnd bool          `json:"logicAnd,omitempty" yaml:"logicAnd"`
	Items    []*AssertItem `json:"items,omitempty" yaml:"items"`
}
type AssertItem struct {
	ValueFrom string      `json:"valueFrom,omitempty" yaml:"valueFrom"`
	Operator  string      `json:"operator,omitempty" yaml:"operator"`
	Value     interface{} `json:"value,omitempty" yaml:"value"`
}

func (c SimpleCondition) Errorf(format string, args ...interface{}) {
	klog.V(3).Info(fmt.Sprintf("SimpleCondition: %s", format), args)
}

func (c SimpleCondition) Pass(req *restful.Request) (ok bool) {

	var extractValue interface{}
	ok = true
	for i := 0; i < len(c.Items); i++ {
		item := c.Items[i]
		if strings.HasPrefix(item.ValueFrom, ValueFromPath) {
			pathKey := strings.TrimSpace(item.ValueFrom[len(ValueFromPath):])
			extractValue = req.PathParameter(pathKey)
		} else if strings.HasPrefix(item.ValueFrom, ValueFromQuery) {
			pathKey := strings.TrimSpace(item.ValueFrom[len(ValueFromQuery):])
			extractValue = req.QueryParameter(pathKey)
		} else if strings.HasPrefix(item.ValueFrom, ValueFromHeader) {
			pathKey := strings.TrimSpace(item.ValueFrom[len(ValueFromHeader):])
			extractValue = req.HeaderParameter(pathKey)
		} else if strings.HasPrefix(item.ValueFrom, ValueFromBody) {
			pathKey := strings.TrimSpace(item.ValueFrom[len(ValueFromBody):])
			extractValue, _ = req.BodyParameter(pathKey)
		}
		result, err := c.validator(extractValue, item.Value, item.Operator)
		if !result && err != nil {
			klog.Warningf("item: %s validator failed, err: %s",
				item.ValueFrom, err.Error())
		}
		if !result {
			ok = false
		} else {
			if !c.LogicAnd {
				return true
			}
			ok = ok == result
			if !ok {
				return false
			}
		}
	}
	return ok
}
func (c SimpleCondition) validator(value, extractValue interface{}, operator string) (success bool, err error) {
	if assertFunc, ok := AssertFunctions[operator]; ok {
		return assertFunc(c, extractValue, value), nil
	} else {
		err = fmt.Errorf("unsupport validator operator: %s", operator)
	}
	return
}
