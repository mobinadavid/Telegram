package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"time"
)

func EmojiSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	emojis := []string{
		"https://example.com/emojis/smile.png",
		"https://example.com/emojis/laugh.png",
		"https://example.com/emojis/heart.png",
		"https://example.com/emojis/sad.png",
		"https://example.com/emojis/angry.png",
	}

	query := `INSERT INTO emojis (emoji_url, created_at, updated_at) VALUES ($1, $2, $3);`

	for _, emojiURL := range emojis {
		_, err := db.Exec(context.Background(), query, emojiURL, time.Now(), time.Now())
		if err != nil {
			log.Printf("Failed to seed emoji %s: %v", emojiURL, err)
		}
	}
	log.Println("Successfully seeded emojis")
}
