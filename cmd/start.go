package main

import (
	"flag"
	"github.com/kube-all/mock-runner/cmd/server"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/klog/v2"
	"os"
)

func commandRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:  "mock-runner",
		Long: `http api mock server`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(2)
		},
	}
	rootCmd.Flags().SortFlags = true
	klog.InitFlags(nil)
	pflag.CommandLine.AddGoFlag(flag.CommandLine.Lookup("v"))
	pflag.CommandLine.AddGoFlag(flag.CommandLine.Lookup("logtostderr"))
	pflag.CommandLine.Set("logtostderr", "true")

	// add sub command
	rootCmd.AddCommand(server.NewServerCommand())
	return rootCmd
}

func main() {
	defer klog.Flush() // flushes all pending log I/O
	command := commandRoot()
	command.Execute()
}
