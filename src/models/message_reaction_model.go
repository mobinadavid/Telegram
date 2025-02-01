package models

import "time"

type MessageReaction struct {
	ID        int64      `json:"id"`
	SenderID  int64      `json:"sender_id"`
	ReactedTo int64      `json:"reacted_to"`
	EmojiID   int64      `json:"emoji_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
