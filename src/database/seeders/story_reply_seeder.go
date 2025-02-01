package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func StoryReplySeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	// Sample story reply data
	storyReplies := []models.StoryReply{
		{Message: "Great story!", OwnerID: 1, StoryID: 1, ChatID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Message: "Amazing content, love it!", OwnerID: 2, StoryID: 2, ChatID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Message: "So inspiring!", OwnerID: 3, StoryID: 2, ChatID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	// SQL query to insert data into story_replies table
	query := `INSERT INTO story_replies (message, owner_id, story_id, chat_id, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6);`

	// Loop through the storyReplies and insert each one into the database
	for _, reply := range storyReplies {
		_, err := db.Exec(context.Background(), query, reply.Message, reply.OwnerID, reply.StoryID, reply.ChatID, reply.CreatedAt, reply.UpdatedAt)
		if err != nil {
			log.Printf("Failed to seed story reply for story ID %d: %v", reply.StoryID, err)
		}
	}

	log.Println("Successfully seeded story replies")
}
