CREATE TABLE IF NOT EXISTS messages (
            id          BIGSERIAL PRIMARY KEY,
            content     TEXT NOT NULL,
            sender_id   BIGINT NOT NULL,
            chat_id   BIGINT NOT NULL,
            created_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
            updated_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
            deleted_at  TIMESTAMP WITH TIME ZONE,

            FOREIGN KEY (sender_id) REFERENCES accounts(id) ON DELETE CASCADE,
          FOREIGN KEY ( chat_id) REFERENCES chats(id) ON DELETE CASCADE
    );

CREATE INDEX IF NOT EXISTS idx_messages_chat_id ON messages(chat_id);


CREATE OR REPLACE FUNCTION update_chat_message_info()
    RETURNS TRIGGER AS $$
BEGIN
    -- Increment the message count for the specific chat
    UPDATE chats
    SET
        messages_count = messages_count + 1,
        last_message_id = NEW.id,
        updated_at = NOW()
    WHERE id = NEW.chat_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;



CREATE TRIGGER trigger_update_chat_message_info
    AFTER INSERT ON messages
    FOR EACH ROW
    EXECUTE FUNCTION update_chat_message_info();



CREATE OR REPLACE FUNCTION decrement_chat_message_info()
    RETURNS TRIGGER AS $$
BEGIN

    UPDATE chats
    SET
        messages_count = messages_count - 1,
        last_message_id = (
            SELECT id FROM messages
            WHERE chat_id = OLD.chat_id
              AND chat_type = OLD.chat_type
            ORDER BY created_at DESC
            LIMIT 1
        ),
        updated_at = NOW()
    WHERE id = OLD.chat_id;

    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_decrement_chat_message_info
    AFTER DELETE ON messages
    FOR EACH ROW
EXECUTE FUNCTION decrement_chat_message_info();

