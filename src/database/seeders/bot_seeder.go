package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func BotSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	// Sample bots to seed
	bots := []models.BotModel{
		{
			Name:             "WeatherBot",
			Description:      "A bot that provides weather updates.",
			OwnerID:          1,
			SubscribersCount: 0,
			Username:         "bot_first",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
		{
			Name:             "NewsBot",
			Description:      "A bot that delivers the latest news.",
			OwnerID:          4,
			SubscribersCount: 0,
			Username:         "bot_second",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
	}

	// SQL query to insert bots
	query := `INSERT INTO bots (name, description, owner_id, subscribers_count, username, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	// Insert each bot
	for _, bot := range bots {
		var botID int64
		err := db.QueryRow(context.Background(), query, bot.Name, bot.Description, bot.OwnerID, bot.SubscribersCount, bot.Username, bot.CreatedAt, bot.UpdatedAt).Scan(&botID)
		if err != nil {
			log.Printf("Failed to seed bot: %s - %v", bot.Name, err)
			continue
		}
	}

	log.Println("Successfully seeded bots")
}
