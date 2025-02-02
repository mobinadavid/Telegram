package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func ChannelSubscriberSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	channelSubscribers := []models.ChannelSubscriber{
		{ChannelID: 3, AccountID: 1, IsAdmin: true, JoinedAt: time.Now(), UpdatedAt: time.Now()},
		{ChannelID: 3, AccountID: 2, IsAdmin: false, JoinedAt: time.Now(), UpdatedAt: time.Now()},
		{ChannelID: 4, AccountID: 3, IsAdmin: false, JoinedAt: time.Now(), UpdatedAt: time.Now()},
		{ChannelID: 4, AccountID: 4, IsAdmin: true, JoinedAt: time.Now(), UpdatedAt: time.Now()},
		{ChannelID: 5, AccountID: 5, IsAdmin: false, JoinedAt: time.Now(), UpdatedAt: time.Now()},
		{ChannelID: 5, AccountID: 6, IsAdmin: true, JoinedAt: time.Now(), UpdatedAt: time.Now()},
		{ChannelID: 3, AccountID: 4, IsAdmin: true, JoinedAt: time.Now(), UpdatedAt: time.Now()},
		{ChannelID: 3, AccountID: 3, IsAdmin: true, JoinedAt: time.Now(), UpdatedAt: time.Now()},
		{ChannelID: 5, AccountID: 3, IsAdmin: true, JoinedAt: time.Now(), UpdatedAt: time.Now()},
	}

	query := `INSERT INTO channel_subscribers (channel_id, account_id, is_admin, joined_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5);`

	for _, channelSubscriber := range channelSubscribers {
		_, err := db.Exec(context.Background(), query, channelSubscriber.ChannelID, channelSubscriber.AccountID, channelSubscriber.IsAdmin, channelSubscriber.JoinedAt, channelSubscriber.UpdatedAt)
		if err != nil {
			log.Printf("Failed to seed channel subscriber for channel ID %d, account ID %d: %v", channelSubscriber.ChannelID, channelSubscriber.AccountID, err)
		}
	}

	log.Println("Successfully seeded channel subscribers.")
}
