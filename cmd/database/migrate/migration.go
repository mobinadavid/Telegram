package migrate

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"telegram/src/database/migrations"
	"telegram/src/database/pgx"
)

// Migrate Commands for interacting with database
var Migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Commands for migrate tables",
}

func init() {
	Migrate.AddCommand(
		migrateUp,
		migrateDown,
	)
}

var migrateUp = &cobra.Command{
	Use:   "up",
	Short: "migrate the tables up",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		beforeMigrate()

		log.Println("Starting migrate up...    ")
		err := migrations.MigrateUp()
		if err != nil {
			log.Fatalf("Failed to migrate up: %v", err)
		}
		log.Println("migrate up successfully.    ")

		afterMigrate()
	},
}

var migrateDown = &cobra.Command{
	Use:   "down",
	Short: "migrate the tables down",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		beforeMigrate()

		log.Println("Starting migrate down...    ")
		err := migrations.MigrateDown()
		if err != nil {
			log.Fatalf("Failed to migrate down: %v.  ", err)
		}
		log.Println("migrate down successfully.    ")

		afterMigrate()
	},
}

func beforeMigrate() {
	// Initialize Env variable
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ENV Service: Failed to  loading .env file. %v.", err)
	}

	// Initialize Database
	err = pgx.Init()
	if err != nil {
		log.Fatalf("Database Service: Failed to Initialize: %v.  ", err)
	}
}

func afterMigrate() {
	err := pgx.Close()
	if err != nil {
		log.Fatalf("Databasde Service: Failed to close database. %v. \n", err)
	}
}
