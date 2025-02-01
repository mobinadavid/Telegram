package models

import "time"

type StoryView struct {
	ID        int64      `json:"id"`
	ViewerID  int64      `json:"viewer_id"`
	StoryID   int64      `json:"story_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
