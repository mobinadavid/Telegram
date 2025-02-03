CREATE TABLE IF NOT EXISTS messages (
     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    content TEXT NOT NULL,
    sender_id UUID NOT NULL,
    chat_id UUID NOT NULL,
    replied_to UUID,
    created_at TIMESTAMPTZ DEFAULT current_timestamp(),
    updated_at TIMESTAMPTZ DEFAULT current_timestamp(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT fk_messages_sender FOREIGN KEY (sender_id) REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_messages_chat FOREIGN KEY (chat_id) REFERENCES chats(id) ON DELETE CASCADE,
    CONSTRAINT fk_messages_reply FOREIGN KEY (replied_to) REFERENCES messages(id) ON DELETE CASCADE
    );