CREATE TABLE IF NOT EXISTS stickers (
      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sticker_url TEXT NOT NULL,
    sticker_pack_id UUID,
    emoji_id UUID,
    created_at TIMESTAMPTZ DEFAULT current_timestamp(),
    updated_at TIMESTAMPTZ DEFAULT current_timestamp(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT fk_sticker_pack FOREIGN KEY (sticker_pack_id) REFERENCES sticker_packs(id) ON DELETE CASCADE,
    CONSTRAINT fk_emoji FOREIGN KEY (emoji_id) REFERENCES emojis(id) ON DELETE CASCADE
    );

CREATE INDEX IF NOT EXISTS idx_stickers_deleted_at
    ON stickers (deleted_at);