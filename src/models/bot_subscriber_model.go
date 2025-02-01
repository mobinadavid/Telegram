package models

import "time"

type BotSubscriber struct {
	BotID     int64     `json:"bot_id"`
	AccountID int64     `json:"account_id"`
	IsAdmin   bool      `json:"is_admin"`
	JoinedAt  time.Time `json:"joined_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
