package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func ChannelSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	channels := []models.Channel{
		{
			Name:        "Tech Talk",
			Description: "Join this community to discuss all things tech!",
			OwnerID:     4,
		},
		{
			Name:        "Foodies Corner",
			Description: "Explore new recipes and food trends with fellow foodies.",
			OwnerID:     2,
		},
		{
			Name:        "Fitness Gurus",
			Description: "Get fit with daily workout tips and motivation.",
			OwnerID:     3,
		},
	}

	query := `INSERT INTO channels (name, description, owner_id, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5);`

	for _, channel := range channels {
		_, err := db.Exec(context.Background(), query, channel.Name, channel.Description, channel.OwnerID, time.Now(), time.Now())
		if err != nil {
			log.Printf("Failed to seed channel %s: %v", channel.Name, err)
		}
	}

	log.Println("Successfully seeded channels")
}
