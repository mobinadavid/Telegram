package migrate

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"telegram/src/database/cockroach"
	"telegram/src/database/migrations"
)

// Migrate Commands for interacting with database
var MigrateCockroach = &cobra.Command{
	Use:   "cockroach",
	Short: "Commands for migrate tables",
}

func init() {
	MigrateCockroach.AddCommand(
		migrateUpCockroach,
		migrateDownCockroach,
	)
}

var migrateUpCockroach = &cobra.Command{
	Use:   "up",
	Short: "migrate the tables up",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		beforeMigrateCockroach()

		log.Println("Starting migrate up...    ")
		err := migrations.MigrateUpCockroach()
		if err != nil {
			log.Fatalf("Failed to migrate up: %v", err)
		}
		log.Println("migrate up successfully.    ")

		afterMigrateCockroach()
	},
}

var migrateDownCockroach = &cobra.Command{
	Use:   "down",
	Short: "migrate the tables down",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		beforeMigrateCockroach()

		log.Println("Starting migrate down...    ")
		err := migrations.MigrateDownCockroach()
		if err != nil {
			log.Fatalf("Failed to migrate down: %v.  ", err)
		}
		log.Println("migrate down successfully.    ")

		afterMigrateCockroach()
	},
}

func beforeMigrateCockroach() {
	// Initialize Env variable
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ENV Service: Failed to  loading .env file. %v.", err)
	}

	// Initialize Database
	err = cockroach.Init()
	if err != nil {
		log.Fatalf("Database Service: Failed to Initialize cockroach: %v.  ", err)
	}
}

func afterMigrateCockroach() {
	err := cockroach.Close()
	if err != nil {
		log.Fatalf("Databasde Service: Failed to close database. %v. \n", err)
	}
}
