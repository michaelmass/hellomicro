package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// VERSION of hellomicro
// nolint:gochecknoglobals
var VERSION string

var rootCmd = &cobra.Command{
	Use:   "hellomicro",
	Short: "hellomicro is a utility server",
	Long:  `A utility server that makes it simpler to test microservices architechture.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("hellomicro version %s\n", VERSION)
	},
}

// Execute the cli
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
