CREATE TABLE IF NOT EXISTS chats (
        chat_id        bigserial PRIMARY KEY,
        account_id     BIGINT NOT NULL,
        type_id        BIGINT NOT NULL,
        chat_type      TEXT NOT NULL CHECK (chat_type IN ('pv', 'channel', 'group', 'bot')),
        is_muted       BOOLEAN DEFAULT FALSE,
        is_archived    BOOLEAN DEFAULT FALSE,
        messages_count BIGINT DEFAULT 0,
        last_message_id BIGINT,
        created_at     TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
        updated_at     TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
        deleted_at     TIMESTAMP WITH TIME ZONE,

      UNIQUE (account_id, type_id, chat_type),
FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE
);

CREATE OR REPLACE FUNCTION validate_chat_type_and_type_id()
    RETURNS TRIGGER AS $$
BEGIN
    -- For 'bot' type, check that type_id exists in the bots table
    IF NEW.chat_type = 'bot' THEN
        IF NOT EXISTS (SELECT 1 FROM bots WHERE id = NEW.type_id) THEN
            RAISE EXCEPTION 'Invalid bot_id for chat type bot';
        END IF;
    END IF;

    -- For 'group' type, check that type_id exists in the groups table
    IF NEW.chat_type = 'group' THEN
        IF NOT EXISTS (SELECT 1 FROM groups WHERE id = NEW.type_id) THEN
            RAISE EXCEPTION 'Invalid group_id for chat type group';
        END IF;
    END IF;

    -- For 'channel' type, check that type_id exists in the channels table
    IF NEW.chat_type = 'channel' THEN
        IF NOT EXISTS (SELECT 1 FROM channels WHERE id = NEW.type_id) THEN
            RAISE EXCEPTION 'Invalid channel_id for chat type channel';
        END IF;
    END IF;

    -- For 'pv' type, no specific check is needed for type_id

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to run before inserting into the chats table
CREATE TRIGGER validate_chat_type_and_type_id_trigger
    BEFORE INSERT ON chats
    FOR EACH ROW
EXECUTE FUNCTION validate_chat_type_and_type_id();