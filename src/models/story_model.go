package models

import "time"

type Story struct {
	ID          int64      `json:"id"`
	OwnerID     int64      `json:"owner_id"`
	ContentType string     `json:"content_type"`
	ContentURL  string     `json:"content_url"`
	Caption     string     `json:"caption"`
	ViewsCount  int64      `json:"views_count"`
	ExpiresAt   time.Time  `json:"expires_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
