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
	"github.com/stretchr/testify/assert"
	"reflect"
)

type AssertFunc func(t assert.TestingT, actual interface{}, expected interface{}, msgAndArgs ...interface{}) bool

var AssertFunctions = map[string]AssertFunc{
	"=":                 assert.EqualValues,
	"eq":                assert.EqualValues,
	"equals":            assert.EqualValues,
	"equal":             assert.EqualValues,
	"<":                 assert.Less,
	"lt":                assert.Less,
	"less_than":         assert.Less,
	"le":                assert.LessOrEqual,
	"<=":                assert.LessOrEqual,
	"less_or_equals":    assert.LessOrEqual,
	"gt":                assert.Greater,
	">":                 assert.Greater,
	"greater_than":      assert.Greater,
	"ge":                assert.GreaterOrEqual,
	"greater_or_equals": assert.GreaterOrEqual,
	">=":                assert.GreaterOrEqual,
	"ne":                assert.NotEqual,
	"!=":                assert.NotEqual,
	"not_equal":         assert.NotEqual,
	"contains":          assert.Contains,
	"type_match":        assert.IsType,
	"len":               lenf,
	"regexMatch":        assert.Regexp,
	"regexNotMatch":     assert.NotRegexp,
	"startWith":         startWith,
}

func startWith(t assert.TestingT, origin interface{}, target interface{}, msgAndArgs ...interface{}) bool {
	if reflect.TypeOf(origin).Kind() != reflect.String || reflect.TypeOf(target).Kind() != reflect.String {
		return false
	}
	return assert.Regexp(t, "$"+origin.(string), target)

}
func lenf(t assert.TestingT, expectedType interface{}, object interface{}, msgAndArgs ...interface{}) bool {
	if reflect.TypeOf(expectedType).Kind() != reflect.Int {
		return false
	}
	return assert.Len(t, object, expectedType.(int), msgAndArgs)
}
