
CREATE TABLE IF NOT EXISTS story_replies (
         id                BIGSERIAL PRIMARY KEY,
         message           Text NOT NULL,
         owner_id          BIGINT,
         story_id           BIGINT,
         chat_id     BIGINT NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at  TIMESTAMP WITH TIME ZONE,

    FOREIGN KEY (owner_id) REFERENCES accounts(id) ON DELETE CASCADE,
    FOREIGN KEY (story_id) REFERENCES stories(id) ON DELETE CASCADE,
         FOREIGN KEY (chat_id) REFERENCES chats(chat_id) ON DELETE CASCADE



);