package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func StickerSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	stickers := []models.Sticker{
		{StickerURL: "https://example.com/stickers/sticker1.png", StickerPackID: 1, EmojiID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{StickerURL: "https://example.com/stickers/sticker2.png", StickerPackID: 1, EmojiID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{StickerURL: "https://example.com/stickers/sticker3.png", StickerPackID: 2, EmojiID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{StickerURL: "https://example.com/stickers/sticker4.png", StickerPackID: 2, EmojiID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{StickerURL: "https://example.com/stickers/sticker5.png", StickerPackID: 3, EmojiID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	query := `INSERT INTO stickers (sticker_url, sticker_pack_id, emoji_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5);`

	for _, sticker := range stickers {
		_, err := db.Exec(context.Background(), query, sticker.StickerURL, sticker.StickerPackID, sticker.EmojiID, sticker.CreatedAt, sticker.UpdatedAt)
		if err != nil {
			log.Printf("Failed to seed sticker %s: %v", sticker.StickerURL, err)
		}
	}
	log.Println("Successfully seeded stickers")
}
