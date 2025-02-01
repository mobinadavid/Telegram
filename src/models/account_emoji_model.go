package models

import "time"

type AccountEmoji struct {
	AccountID  int64      `json:"account_id"`
	EmojiID    int64      `json:"emoji_id"`
	UsedAt     *time.Time `json:"used_at,omitempty"`
	UsageCount int64      `json:"usage_count"`
}
