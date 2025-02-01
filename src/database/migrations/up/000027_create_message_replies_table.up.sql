CREATE TABLE IF NOT EXISTS message_replies (
        id          BIGSERIAL PRIMARY KEY,
        content     TEXT NOT NULL,
        sender_id   BIGINT NOT NULL,
        reply_to   BIGINT NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at  TIMESTAMP WITH TIME ZONE,
        FOREIGN KEY (reply_to) REFERENCES messages(id) ON DELETE CASCADE
    );

CREATE OR REPLACE FUNCTION update_chat_message_reply_info()
    RETURNS TRIGGER AS $$
BEGIN
    UPDATE chats
    SET
        messages_count = messages_count + 1,
        last_message_id = NEW.id,
        updated_at = NOW()
    WHERE id = (
        SELECT chat_id FROM messages WHERE id = NEW.reply_to
    );

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION decrement_chat_message_reply_info()
    RETURNS TRIGGER AS $$
BEGIN
    UPDATE chats
    SET
        messages_count = messages_count - 1,
        last_message_id = (
            SELECT id FROM message_replies
            WHERE reply_to = OLD.reply_to
            ORDER BY created_at DESC
            LIMIT 1
        ),
        updated_at = NOW()
    WHERE id = (
        SELECT chat_id FROM messages WHERE id = OLD.reply_to
    );

    RETURN OLD;
END;
$$ LANGUAGE plpgsql;



CREATE TRIGGER trigger_update_chat_message_reply_info
    AFTER INSERT ON message_replies
    FOR EACH ROW
EXECUTE FUNCTION update_chat_message_reply_info();

CREATE TRIGGER trigger_decrement_chat_message_reply_info
    AFTER DELETE ON message_replies
    FOR EACH ROW
EXECUTE FUNCTION decrement_chat_message_reply_info();


