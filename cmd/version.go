package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newVersion())
}

func newVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of hellomicro",
		Long:  `All software has versions. This is hellomicro's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("hellomicro v%s\n", VERSION)
		},
	}
}
