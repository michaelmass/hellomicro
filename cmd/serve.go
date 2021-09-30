package cmd

import (
	"github.com/michaelmass/hellomicro/pkg/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newServe())
}

func newServe() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the hellomicro server",
		Long:  `Start the hellomicro server`,
		Run: func(cmd *cobra.Command, args []string) {
			server.Run()
		},
	}
}
