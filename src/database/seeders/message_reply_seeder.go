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
		{Context: "I agree with you!", SenderID: 2, ReplyTo: 1},    // Replying to message 1
		{Context: "I'll be there!", SenderID: 3, ReplyTo: 2},       // Replying to message 2
		{Context: "That's awesome!", SenderID: 4, ReplyTo: 3},      // Replying to message 3
		{Context: "Let's meet tomorrow!", SenderID: 5, ReplyTo: 4}, // Replying to message 4
		{Context: "Got it, thanks!", SenderID: 6, ReplyTo: 5},      // Replying to message 5
	}

	query := `INSERT INTO message_replies (context, sender_id, reply_to, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5);`

	for _, reply := range messageReplies {
		_, err := db.Exec(context.Background(), query, reply.Context, reply.SenderID, reply.ReplyTo, time.Now(), time.Now())
		if err != nil {
			log.Printf("Failed to seed message reply to message %d: %v", reply.ReplyTo, err)
		}
	}

	log.Printf("Successfully seeded message replies")
}
