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

package scaffold

import (
	"fmt"
	"github.com/kube-all/mock-runner/cmd/scaffold/options"
	"github.com/kube-all/mock-runner/pkg/embeds"
	"github.com/spf13/cobra"
	"io/ioutil"
	"k8s.io/klog/v2"
	"os"
	"path"
	"strings"
)

func NewScaffoldCommand() *cobra.Command {
	s := options.NewScaffoldOptions()
	cmd := &cobra.Command{
		Use: "new",
		Long: `mock new scaffold
mock new ${project}
`,
		Short:        "mock new",
		Example:      "mock new mock-project",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("you must specify a project name")
			}
			s.Project = args[0]
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true
			return run(s)
		},
	}
	flags := cmd.Flags()
	flags.StringVarP(&s.Project, "project", "p", "", "mock repo path")
	return cmd
}
func run(o *options.ScaffoldOptions) (err error) {
	basePath := "templates"
	dirs, files, err := getAllFiles(basePath)
	if err != nil {
		klog.Fatalf("get mock template failed, err: %s", err.Error())
	}
	// create dirs
	for _, d := range dirs {
		cdir := strings.Replace(d, basePath, o.Project, 1)
		err = createDir(cdir)
		if err != nil {
			klog.Fatalf("create dir: %s failed, err: %s", cdir, err.Error())
		}
	}
	// create files
	for _, f := range files {
		if strings.HasSuffix(f, "gitignore") {
			continue
		}
		content, err := embeds.ProjectTemplate.ReadFile(f)
		if err != nil {
			klog.Fatalf("read file : %s failed, err: %s", f, err.Error())
		}
		fName := strings.Replace(f, basePath, o.Project, 1)
		err = createFile(fName, content, true)
		if err != nil {
			klog.Fatalf("write file : %s failed, err: %s", fName, err.Error())
		}
	}
	// 写入忽略文件
	if content, er := embeds.ProjectTemplate.ReadFile(path.Join(basePath, "gitignore")); er == nil {
		createFile(path.Join(o.Project, ".gitignore"), content, true)
	}
	return nil
}
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func createDir(dir string) (err error) {
	if exist, _ := pathExists(dir); !exist {
		err = os.MkdirAll(dir, os.ModePerm)
	}
	return err
}
func createFile(fName string, content []byte, cover bool) (err error) {
	if e, _ := pathExists(fName); e {
		if !cover {
			return fmt.Errorf("file: %s exist", fName)
		} else {
			klog.Infof("file: %s exist, will try to cover it's content", fName)
			err = os.Remove(fName)
			if err == nil {
				return ioutil.WriteFile(fName, content, os.ModePerm)
			} else {
				klog.Infof("file: %s exist, delete it failed, err: %s", fName, err.Error())
				return err
			}
		}
	} else {
		return ioutil.WriteFile(fName, content, os.ModePerm)
	}

}
func getAllFiles(dirPath string) (dirs []string, files []string, err error) {
	fs, err := embeds.ProjectTemplate.ReadDir(dirPath)
	if err != nil {
		return
	}
	for _, f := range fs {
		if f.IsDir() {
			dirs = append(dirs, path.Join(dirPath, f.Name()))
			ds, fs, er := getAllFiles(path.Join(dirPath, f.Name()))
			if er == nil {
				dirs = append(dirs, ds...)
				files = append(files, fs...)
			}
		} else {
			files = append(files, path.Join(dirPath, f.Name()))
		}
	}
	return
}
