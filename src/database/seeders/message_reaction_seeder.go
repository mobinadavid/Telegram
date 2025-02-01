package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func MessageReactionSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	// Seed data for message reactions
	messageReactions := []models.MessageReaction{
		{SenderID: 2, ReactedTo: 1, EmojiID: 1},
		{SenderID: 3, ReactedTo: 2, EmojiID: 2},
		{SenderID: 4, ReactedTo: 3, EmojiID: 3},
		{SenderID: 5, ReactedTo: 4, EmojiID: 4},
		{SenderID: 6, ReactedTo: 5, EmojiID: 5},
	}

	// The query to insert the data
	query := `INSERT INTO message_reactions (sender_id, reacted_to, emoji_id, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5);`

	// Loop through each message reaction and insert it
	for _, reaction := range messageReactions {
		_, err := db.Exec(context.Background(), query, reaction.SenderID, reaction.ReactedTo, reaction.EmojiID, time.Now(), time.Now())
		if err != nil {
			log.Printf("Failed to seed message reaction to message %d: %v", reaction.ReactedTo, err)
		}
	}

	log.Printf("Successfully seeded message reactions")
}
