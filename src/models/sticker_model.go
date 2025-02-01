package models

import "time"

type Sticker struct {
	ID            int64      `json:"id"`
	StickerURL    string     `json:"sticker_url"`
	StickerPackID int64      `json:"sticker_pack_id"`
	EmojiID       int64      `json:"emoji_id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
}
