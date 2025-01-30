CREATE TABLE IF NOT EXISTS message_replies (
        id          BIGSERIAL PRIMARY KEY,
        context     TEXT NOT NULL,
        sender_id   BIGINT NOT NULL,
        reply_to   BIGINT NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at  TIMESTAMP WITH TIME ZONE,
        FOREIGN KEY (reply_to) REFERENCES messages(id) ON DELETE CASCADE
    );
