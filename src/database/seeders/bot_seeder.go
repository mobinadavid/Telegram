package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func BotSeeder() {
	// Get the database instance
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	// Define some sample bots to seed
	bots := []models.Bot{
		{
			Name:        "BotOne",
			Description: "This is Bot One description.",
			OwnerID:     1, // Assuming the account with ID 1 is the owner
			Username:    "botone",
		},
		{
			Name:        "BotTwo",
			Description: "This is Bot Two description.",
			OwnerID:     2, // Assuming the account with ID 2 is the owner
			Username:    "bottwo",
		},
	}

	// Define the SQL query for inserting bots
	query := `INSERT INTO bots (name, description, owner_id, username, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6);`

	// Loop over the bot data and insert each bot into the database
	for _, bot := range bots {
		_, err := db.Exec(context.Background(), query, bot.Name, bot.Description, bot.OwnerID, bot.Username, time.Now(), time.Now())
		if err != nil {
			log.Printf("Failed to seed bot %s: %v", bot.Name, err)
		}
	}

	log.Printf("Successfully seeded bots")
}
