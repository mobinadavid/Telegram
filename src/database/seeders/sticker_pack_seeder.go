package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"time"
)

func StickerPackSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	stickerPacks := []string{
		"Funny Stickers",
		"Animal Stickers",
		"Meme Stickers",
		"Emoji Stickers",
		"Cartoon Stickers",
	}

	query := `INSERT INTO sticker_packs (name, created_at, updated_at) VALUES ($1, $2, $3);`

	for _, name := range stickerPacks {
		_, err := db.Exec(context.Background(), query, name, time.Now(), time.Now())
		if err != nil {
			log.Printf("Failed to seed sticker pack %s: %v", name, err)
		}
	}
	log.Println("Successfully seeded sticker packs")
}
