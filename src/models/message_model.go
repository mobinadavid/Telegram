package models

import "time"

type Message struct {
	ID        int64      `json:"id"`
	Content   string     `json:"Content"`
	SenderID  int64      `json:"sender_id"`
	ChatID    int64      `json:"chat_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
