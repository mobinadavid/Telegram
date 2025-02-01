package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func AccountStickerSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	accountStickers := []models.AccountSticker{
		{AccountID: 3, StickerID: 1, UsedAt: timePtr(time.Now().Add(-24 * time.Hour)), UsageCount: 10},
		{AccountID: 3, StickerID: 2, UsedAt: timePtr(time.Now().Add(-48 * time.Hour)), UsageCount: 5},
		{AccountID: 3, StickerID: 3, UsedAt: timePtr(time.Now().Add(-72 * time.Hour)), UsageCount: 8},
		{AccountID: 4, StickerID: 4, UsedAt: timePtr(time.Now().Add(-96 * time.Hour)), UsageCount: 3},
		{AccountID: 3, StickerID: 5, UsedAt: timePtr(time.Now().Add(-120 * time.Hour)), UsageCount: 6},
		{AccountID: 1, StickerID: 1, UsedAt: timePtr(time.Now().Add(-24 * time.Hour)), UsageCount: 10},
		{AccountID: 2, StickerID: 2, UsedAt: timePtr(time.Now().Add(-48 * time.Hour)), UsageCount: 5},
		{AccountID: 1, StickerID: 3, UsedAt: timePtr(time.Now().Add(-72 * time.Hour)), UsageCount: 8},
	}

	query := `INSERT INTO account_stickers (account_id, sticker_id, used_at, usage_count) VALUES ($1, $2, $3, $4);`

	for _, as := range accountStickers {
		_, err := db.Exec(context.Background(), query, as.AccountID, as.StickerID, as.UsedAt, as.UsageCount)
		if err != nil {
			log.Printf("Failed to seed account_sticker with account_id %d and sticker_id %d: %v", as.AccountID, as.StickerID, err)
		}
	}
	log.Println("Successfully seeded account_stickers")
}

func timePtr(t time.Time) *time.Time {
	return &t
}
