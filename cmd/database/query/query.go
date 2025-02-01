package query

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"telegram/src/database/pgx"
	"time"
)

// Query Commands for interacting with database
var Query = &cobra.Command{
	Use:   "query",
	Short: "Commands for seeders tables",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		beforeMigrate()

		// add seeders
		log.Println("Starting querying table...  ")

		log.Println("seeding tables successfully. ")

		afterMigrate()
	},
}

func beforeMigrate() {
	// Initialize Env variable
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ENV Service: Failed to  loading .env file. %v.    timestamp: %s", err, time.Now().String())
	}

	// Initialize Database
	err = pgx.Init()
	if err != nil {
		log.Fatalf("Database Service: Failed to Initialize: %v.    timestamp: %s", err, time.Now().String())
	}
}

func afterMigrate() {
	err := pgx.Close()
	if err != nil {
		log.Fatalf("Databasde Service: Failed to close database. %v.    timestamp: %s \n", err, time.Now().String())
	}
}
