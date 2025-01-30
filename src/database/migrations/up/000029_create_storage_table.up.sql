CREATE TYPE file_type_enum AS ENUM ('image', 'video', 'audio', 'document', 'other');

CREATE TABLE IF NOT EXISTS storage (
    id          BIGSERIAL PRIMARY KEY,
    file_name     TEXT ,
    file_path     Text not null,
    file_size     bigint,
    file_type     file_type_enum,
    mime_type     varchar(255),
    uploader_id   BIGINT NOT NULL,
    message_id   BIGINT ,
    message_reply_id BIGINT ,
    chat_id     BIGINT NOT NULL,
    chat_type   TEXT NOT NULL CHECK (chat_type IN ('account','channel', 'group', 'bot')),
    uploaded_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at  TIMESTAMP WITH TIME ZONE,

    FOREIGN KEY (uploader_id) REFERENCES accounts(id) ON DELETE CASCADE,
    FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE,
    FOREIGN KEY (message_reply_id) REFERENCES message_replies(id) ON DELETE CASCADE,
    FOREIGN KEY (uploader_id, chat_id, chat_type) REFERENCES chats(account_id, chat_id, chat_type) ON DELETE CASCADE
    );

CREATE INDEX IF NOT EXISTS idx_storage_chat ON storage(chat_id, chat_type);
