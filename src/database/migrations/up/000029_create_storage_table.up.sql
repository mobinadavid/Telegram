CREATE TYPE file_type_enum AS ENUM ('image', 'video', 'audio', 'document', 'other');

CREATE TABLE IF NOT EXISTS storage (
    id          BIGSERIAL PRIMARY KEY,
    file_name     TEXT ,
    file_path     Text not null,
    file_size     bigint,
    file_type     file_type_enum not null ,
    mime_type     varchar(255) not null ,
    uploader_id   BIGINT NOT NULL,
    message_id   BIGINT ,
    message_reply_id BIGINT ,
    chat_id     BIGINT NOT NULL,
    uploaded_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at  TIMESTAMP WITH TIME ZONE,

    FOREIGN KEY (uploader_id) REFERENCES accounts(id) ON DELETE CASCADE,
    FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE,
    FOREIGN KEY (message_reply_id) REFERENCES message_replies(id) ON DELETE CASCADE,
    FOREIGN KEY (chat_id) REFERENCES chats(id) ON DELETE CASCADE
    );

CREATE INDEX IF NOT EXISTS idx_storage_chat ON storage(chat_id);
