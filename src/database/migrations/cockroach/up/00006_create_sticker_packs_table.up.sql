CREATE TABLE IF NOT EXISTS sticker_packs (
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT current_timestamp(),
    updated_at TIMESTAMPTZ DEFAULT current_timestamp(),
    deleted_at TIMESTAMPTZ
    );