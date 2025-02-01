CREATE TABLE IF NOT EXISTS chats (
        id       bigserial PRIMARY KEY,
        chat_type      TEXT NOT NULL CHECK (chat_type IN ('pv', 'channel', 'group', 'bot')),
        messages_count BIGINT DEFAULT 0,
        last_message_id BIGINT,
        created_at     TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
        updated_at     TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
        deleted_at     TIMESTAMP WITH TIME ZONE
);
