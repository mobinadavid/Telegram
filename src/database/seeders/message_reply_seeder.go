package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func MessageReplySeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	messageReplies := []models.MessageReply{
		{Content: "I agree with you!", SenderID: 2, ReplyTo: 1},    // Replying to message 1
		{Content: "I'll be there!", SenderID: 3, ReplyTo: 2},       // Replying to message 2
		{Content: "That's awesome!", SenderID: 4, ReplyTo: 3},      // Replying to message 3
		{Content: "Let's meet tomorrow!", SenderID: 5, ReplyTo: 4}, // Replying to message 4
		{Content: "Got it, thanks!", SenderID: 6, ReplyTo: 5},      // Replying to message 5
	}

	query := `INSERT INTO message_replies (content, sender_id, reply_to, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5);`

	for _, reply := range messageReplies {
		_, err := db.Exec(context.Background(), query, reply.Content, reply.SenderID, reply.ReplyTo, time.Now(), time.Now())
		if err != nil {
			log.Printf("Failed to seed message reply to message %d: %v", reply.ReplyTo, err)
		}
	}

	log.Printf("Successfully seeded message replies")
}
