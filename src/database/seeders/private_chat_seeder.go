package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

// PrivateChatSeeder will seed the private_chats table.
func PrivateChatSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	// Sample private chats to seed (accounts 1 and 2 are the participants)
	privateChats := []models.PrivateChatModel{
		{
			FirstAccountID:  1,
			SecondAccountID: 2,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		{
			FirstAccountID:  3,
			SecondAccountID: 4,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
	}

	// SQL query to insert private chat
	query := `INSERT INTO private_chats (first_account_id, second_account_id, created_at, updated_at)
	          VALUES ($1, $2, $3, $4) RETURNING chat_id`

	// Insert each private chat and capture the chat_id
	for _, privateChat := range privateChats {
		var chatID int64
		err := db.QueryRow(context.Background(), query, privateChat.FirstAccountID, privateChat.SecondAccountID, privateChat.CreatedAt, privateChat.UpdatedAt).Scan(&chatID)
		if err != nil {
			log.Printf("Failed to seed private chat between account ID %d and account ID %d: %v", privateChat.FirstAccountID, privateChat.SecondAccountID, err)
			continue
		}
	}

	log.Println("Successfully seeded private chats")
}
