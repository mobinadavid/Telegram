package models

import "time"

type MessageReply struct {
	ID        int64      `json:"id"`
	Context   string     `json:"context"`
	SenderID  int64      `json:"sender_id"`
	ReplyTo   int64      `json:"reply_to"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
