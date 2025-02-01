package models

import "time"

type ProfilePicture struct {
	ID         int64      `json:"id"`
	PictureURL string     `json:"picture_url"`
	OwnerID    int64      `json:"owner_id"`
	OwnerType  string     `json:"owner_type"` // "account", "group", or "channel"
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}
