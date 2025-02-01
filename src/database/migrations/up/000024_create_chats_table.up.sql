CREATE TABLE IF NOT EXISTS chats (
account_id      BIGINT NOT NULL,
chat_id         BIGINT NOT NULL,
chat_type       TEXT NOT NULL CHECK (chat_type IN ('account', 'channel', 'group', 'bot')),
    is_muted        BOOLEAN DEFAULT FALSE,
    is_archived     BOOLEAN DEFAULT FALSE,
    messages_count  BIGINT DEFAULT 0,
    last_message_id BIGINT,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at      TIMESTAMP WITH TIME ZONE,

                                  PRIMARY KEY (account_id, chat_id, chat_type),
    FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE
    );
