package server

import (
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	"net/http"
)

func NewServerCommand() *cobra.Command {
	s := NewServerOptions()
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
func run(o *Options) (err error) {

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		klog.Errorf("start http mock server failed, err: %s", err.Error())
	}

	klog.Info("mock server shutting down")
	return err
}
