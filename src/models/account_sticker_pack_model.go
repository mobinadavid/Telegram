package models

import "time"

type AccountStickerPack struct {
	AccountID     int64      `json:"account_id"`
	StickerPackID int64      `json:"sticker_pack_id"`
	UsedAt        *time.Time `json:"used_at,omitempty"`
	UsageCount    int64      `json:"usage_count"`
}
