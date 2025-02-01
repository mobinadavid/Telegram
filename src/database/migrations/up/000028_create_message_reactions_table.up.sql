CREATE TABLE IF NOT EXISTS message_reactions (
        id          BIGSERIAL PRIMARY KEY,
        sender_id   BIGINT NOT NULL,
        reacted_to   BIGINT NOT NULL,
        emoji_id     BIGINT NOT NULL,
       created_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
       updated_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
       deleted_at  TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (reacted_to) REFERENCES messages(id) ON DELETE CASCADE,
    FOREIGN KEY (emoji_id) REFERENCES emojis(id) ON DELETE CASCADE

    );