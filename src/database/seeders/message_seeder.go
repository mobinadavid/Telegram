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
	messageID1 := int64(1)
	messageID3 := int64(3)

	messages := []models.Message{
		{Content: "Hello, how are you?", SenderID: 3, ChatID: 10, RepliedTo: nil},
		{Content: "I'm good, thanks for asking!", SenderID: 4, ChatID: 10, RepliedTo: &messageID1},
		{Content: "Are you coming to the party tomorrow?", SenderID: 2, ChatID: 6, RepliedTo: nil},
		{Content: "Yes, I will be there.", SenderID: 3, ChatID: 7, RepliedTo: &messageID3},
		{Content: "Hey, don't forget the meeting at 3 PM.", SenderID: 1, ChatID: 6, RepliedTo: nil},
		{Content: "Got it, see you then!", SenderID: 2, ChatID: 7, RepliedTo: nil},
		{Content: "see you!", SenderID: 2, ChatID: 1, RepliedTo: nil},
		{Content: "byee!", SenderID: 2, ChatID: 1, RepliedTo: nil},
	}

	query := `INSERT INTO messages (content, sender_id, chat_id,replied_to, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4,$5, $6);`

	for _, message := range messages {
		_, err := db.Exec(context.Background(), query, message.Content, message.SenderID, message.ChatID, message.RepliedTo, time.Now(), time.Now())
		if err != nil {
			log.Printf("Failed to seed message %s: %v", message.Content, err)
		}
	}

	log.Printf("Successfully seeded messages")
}
