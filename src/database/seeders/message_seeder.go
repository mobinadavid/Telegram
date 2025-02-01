package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func MessageSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	// Sample data for seeding messages
	messages := []models.Message{
		{Content: "Hello, how are you?", SenderID: 3, ChatID: 10},
		{Content: "I'm good, thanks for asking!", SenderID: 4, ChatID: 10},
		{Content: "Are you coming to the party tomorrow?", SenderID: 2, ChatID: 7},
		{Content: "Yes, I will be there.", SenderID: 3, ChatID: 7},
		{Content: "Hey, don't forget the meeting at 3 PM.", SenderID: 1, ChatID: 9},
		{Content: "Got it, see you then!", SenderID: 2, ChatID: 9},
	}

	query := `INSERT INTO messages (content, sender_id, chat_id, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5);`

	for _, message := range messages {
		_, err := db.Exec(context.Background(), query, message.Content, message.SenderID, message.ChatID, time.Now(), time.Now())
		if err != nil {
			log.Printf("Failed to seed message %s: %v", message.Content, err)
		}
	}

	log.Printf("Successfully seeded messages")
}
