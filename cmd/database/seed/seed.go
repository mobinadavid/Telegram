package seed

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"telegram/src/database/pgx"
	"time"
)

// Seed Commands for interacting with database
var Seed = &cobra.Command{
	Use:   "seeders",
	Short: "Commands for seeders tables",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		beforeMigrate()

		// add seeders
		log.Println("Starting seeding table...    " + "timestamp" + time.Now().String())
		//seeders.ProfileSeeder()
		//seeders.CustomerSeeder()
		//seeders.EmployeeSeeder()
		//seeders.AccountSeeder()
		//seeders.TransactionSeeder()
		//seeders.LoanSeeder()
		//seeders.InstallmentSeeder()

		log.Println("seeding table successfully.    " + "timestamp" + time.Now().String())

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
