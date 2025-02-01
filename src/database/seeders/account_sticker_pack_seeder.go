package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func AccountStickerPackSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	accountStickerPacks := []models.AccountStickerPack{
		{AccountID: 1, StickerPackID: 1, UsedAt: timePtr(time.Now().Add(-24 * time.Hour)), UsageCount: 15},
		{AccountID: 2, StickerPackID: 2, UsedAt: timePtr(time.Now().Add(-48 * time.Hour)), UsageCount: 10},
		{AccountID: 3, StickerPackID: 3, UsedAt: timePtr(time.Now().Add(-72 * time.Hour)), UsageCount: 7},
		{AccountID: 4, StickerPackID: 1, UsedAt: timePtr(time.Now().Add(-96 * time.Hour)), UsageCount: 12},
		{AccountID: 5, StickerPackID: 2, UsedAt: timePtr(time.Now().Add(-120 * time.Hour)), UsageCount: 9},
	}

	query := `INSERT INTO account_sticker_packs (account_id, sticker_pack_id, used_at, usage_count) VALUES ($1, $2, $3, $4);`

	for _, asp := range accountStickerPacks {
		_, err := db.Exec(context.Background(), query, asp.AccountID, asp.StickerPackID, asp.UsedAt, asp.UsageCount)
		if err != nil {
			log.Printf("Failed to seed account_sticker_pack with account_id %d and sticker_pack_id %d: %v", asp.AccountID, asp.StickerPackID, err)
		}
	}
	log.Println("Successfully seeded account_sticker_packs")
}
