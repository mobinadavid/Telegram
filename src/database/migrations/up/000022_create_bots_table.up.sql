create table if not exists bots
(
    id              bigserial primary key ,
    name                 VARCHAR(255) NOT NULL,
    description          VARCHAR(255),
    owner_id             bigint NOT NULL,
    subscribers_count    bigint default 0,
    username             VARCHAR(255),
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone,
       FOREIGN KEY (owner_id) REFERENCES accounts (id) ON DELETE CASCADE,
    FOREIGN KEY (id) REFERENCES chats (id) ON DELETE CASCADE
    );

-- Create function to insert chat before inserting bot
CREATE OR REPLACE FUNCTION create_chat_for_bot()
    RETURNS TRIGGER AS $$
DECLARE
    new_chat_id BIGINT;
BEGIN
    -- Step 1: Insert a new chat record for the bot
    INSERT INTO chats (chat_type, created_at, updated_at)
    VALUES ('bot', NOW(), NOW())
    RETURNING id INTO new_chat_id;

    -- Step 2: Set the newly created chat_id to the bot
    NEW.id := new_chat_id;

    -- Return the new bot record with the chat_id set
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger to execute the function after inserting a new bot
CREATE TRIGGER create_chat_for_new_bot
    BEFORE INSERT ON bots
    FOR EACH ROW
EXECUTE FUNCTION create_chat_for_bot();