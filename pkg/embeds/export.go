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

package embeds

import (
	"embed"
	"io/fs"
	"k8s.io/klog/v2"
	"net/http"
)

//go:embed swagger-ui
var staticFiles embed.FS

func StaticFileSystem() http.FileSystem {

	fsys, err := fs.Sub(staticFiles, "swagger-ui")
	if err != nil {
		klog.Fatal(err)
	}
	return http.FS(fsys)
}
