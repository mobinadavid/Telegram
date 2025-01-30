package app

import (
	"github.com/spf13/cobra"
	"log"
	"telegram/src/bootstrap"
)

// App Commands for interacting with apps
var AppCmd = &cobra.Command{
	Use:   "app",
	Short: "Commands for interacting with apps.",
}

func init() {
	AppCmd.AddCommand(
		bootstrapCmd,
	)
}

// bootstrapCmd is sub command for App that bootstrap the application
// usage: go run main.go app bootstrap or ./binary app bootstrap
var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstraps the application and it's related services.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting the Application...    ")
		bootstrap.Init()
		log.Println("Application exited successfully.    ")
	},
}
