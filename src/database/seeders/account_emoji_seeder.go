package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
)

func AccountEmojiSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	accountEmojis := []models.AccountEmoji{
		{AccountID: 1, EmojiID: 1, UsedAt: nil, UsageCount: 5},
		{AccountID: 2, EmojiID: 2, UsedAt: nil, UsageCount: 3},
		{AccountID: 3, EmojiID: 3, UsedAt: nil, UsageCount: 7},
		{AccountID: 4, EmojiID: 4, UsedAt: nil, UsageCount: 2},
		{AccountID: 5, EmojiID: 5, UsedAt: nil, UsageCount: 4},
	}

	query := `INSERT INTO account_emoji (account_id, emoji_id, used_at, usage_count) VALUES ($1, $2, $3, $4);`

	for _, ae := range accountEmojis {
		_, err := db.Exec(context.Background(), query, ae.AccountID, ae.EmojiID, ae.UsedAt, ae.UsageCount)
		if err != nil {
			log.Printf("Failed to seed account_emoji with account_id %d and emoji_id %d: %v", ae.AccountID, ae.EmojiID, err)
		}
	}
	log.Println("Successfully seeded account_emoji")
}
