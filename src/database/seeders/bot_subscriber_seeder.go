package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func BotSubscriberSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	botSubscribers := []models.BotSubscriber{
		{BotID: 1, AccountID: 1, IsAdmin: true, JoinedAt: time.Now(), UpdatedAt: time.Now()},
		{BotID: 1, AccountID: 2, IsAdmin: false, JoinedAt: time.Now(), UpdatedAt: time.Now()},
		{BotID: 2, AccountID: 3, IsAdmin: false, JoinedAt: time.Now(), UpdatedAt: time.Now()},
		{BotID: 2, AccountID: 4, IsAdmin: true, JoinedAt: time.Now(), UpdatedAt: time.Now()},
		{BotID: 1, AccountID: 5, IsAdmin: false, JoinedAt: time.Now(), UpdatedAt: time.Now()},
		{BotID: 1, AccountID: 6, IsAdmin: true, JoinedAt: time.Now(), UpdatedAt: time.Now()},
	}

	query := `INSERT INTO bot_subscribers (bot_id, account_id, is_admin, joined_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5);`

	for _, botSubscriber := range botSubscribers {
		_, err := db.Exec(context.Background(), query, botSubscriber.BotID, botSubscriber.AccountID, botSubscriber.IsAdmin, botSubscriber.JoinedAt, botSubscriber.UpdatedAt)
		if err != nil {
			log.Printf("Failed to seed bot subscriber for bot ID %d, account ID %d: %v", botSubscriber.BotID, botSubscriber.AccountID, err)
		}
	}

	log.Println("Successfully seeded bot subscribers.")
}
