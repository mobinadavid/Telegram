package models

import "time"

type StoryReply struct {
	ID        int64      `json:"id"`
	Message   string     `json:"message"`
	OwnerID   int64      `json:"owner_id"`
	StoryID   int64      `json:"story_id"`
	ChatID    int64      `json:"chat_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
