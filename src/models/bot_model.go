package models

import "time"

type BotModel struct {
	Name             string
	Description      string
	OwnerID          int64
	SubscribersCount int64
	Username         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
