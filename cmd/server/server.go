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
package server

import (
	"github.com/kube-all/mock-runner/cmd/server/options"
	"github.com/kube-all/mock-runner/pkg/api"
	"github.com/spf13/cobra"
)

func NewServerCommand() *cobra.Command {
	s := options.NewServerOptions()
	cmd := &cobra.Command{
		Use:          "server",
		Long:         `mock`,
		Short:        "mock server ",
		Example:      "mock server -p={mockdir}",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				s.Path = args[0]
			}
			return run(s)
		},
	}
	flags := cmd.Flags()
	flags.StringVarP(&s.Path, "path", "p", "", "mock repo path")
	return cmd
}

// run will load mock api and start http server
func run(o *options.Options) (err error) {
	api.AddResource(o)
	return err
}
