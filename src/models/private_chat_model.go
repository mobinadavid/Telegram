package models

import "time"

type PrivateChatModel struct {
	ChatID          int64     `json:"chat_id"`           // Primary key for the chat
	FirstAccountID  int64     `json:"first_account_id"`  // ID of the first account in the chat
	SecondAccountID int64     `json:"second_account_id"` // ID of the second account in the chat
	CreatedAt       time.Time `json:"created_at"`        // Timestamp for when the chat was created
	UpdatedAt       time.Time `json:"updated_at"`        // Timestamp for when the chat was last updated
	DeletedAt       time.Time `json:"deleted_at"`        // Timestamp for when the chat was deleted (soft delete)
}
