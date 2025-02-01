package models

import "time"

type AccountSticker struct {
	AccountID  int64      `json:"account_id"`
	StickerID  int64      `json:"sticker_id"`
	UsedAt     *time.Time `json:"used_at,omitempty"`
	UsageCount int64      `json:"usage_count"`
}
