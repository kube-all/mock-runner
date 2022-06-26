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

package utils

import (
	"io/fs"
	"io/ioutil"
	"k8s.io/klog/v2"
	"path/filepath"
)

func LoadDirFileData(dataPath string) map[string][]byte {
	dataMap := make(map[string][]byte)
	filepath.Walk(dataPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			klog.Warningf("load data from path: %s failed, err: %s", path, err.Error())
			return err
		}
		if info.IsDir() {
			return nil
		}
		relPath, err := filepath.Rel(dataPath, path)
		if err != nil {
			klog.Warningf("file path: %s rel: %s, err: %s ", path, dataPath, err.Error())
			return nil
		}
		if data, err := ioutil.ReadFile(filepath.Join(dataPath, relPath)); err == nil {
			dataMap[relPath] = data
		} else {
			klog.Warningf("read file: %s failed, err: %s", relPath, err.Error())
		}
		return nil
	})
	return dataMap
}
