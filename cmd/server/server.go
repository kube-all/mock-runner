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
		Long:         `mock-runner`,
		Short:        "mock-runner server ",
		Example:      "mock-runner server -p={mockdir}",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(s)
		},
	}
	return cmd
}

// run will load mock api and start http server
func run(o *options.Options) (err error) {
	api.AddResource(o)
	return err
}
