CREATE TABLE IF NOT EXISTS emojis (
     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    emoji_url TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT current_timestamp(),
    updated_at TIMESTAMPTZ DEFAULT current_timestamp(),
    deleted_at TIMESTAMPTZ
    );