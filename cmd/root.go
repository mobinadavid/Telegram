package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"telegram/cmd/app"
	"telegram/cmd/database"
	"telegram/src/database/config"
)

var rootCmd = &cobra.Command{
	Use:   "db-project",
	Short: "db-project CLI",
	Long:  "Golang db-project CLI",
}

func init() {
	cobra.OnInitialize(
		initConfig,
	)

	rootCmd.AddCommand(
		app.AppCmd,
		database.DatabaseCmd,
	)
}

func Execute() error {
	return rootCmd.Execute()
}

func initConfig() {
	// Initialize configuration
	err := config.Init()
	if err != nil {
		log.Fatalf("Config Service: Failed to Initialize. %v", err)
	}
	log.Println("Config Service: Initialized Successfully.")
}
