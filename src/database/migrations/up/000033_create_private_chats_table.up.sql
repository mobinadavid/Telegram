CREATE TABLE IF NOT EXISTS private_chats
(
    chat_id       BIGSERIAL primary Key ,
    first_account_id     BIGINT NOT NULL,
    second_account_id   BIGINT NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE,
    updated_at    TIMESTAMP WITH TIME ZONE,
    deleted_at    TIMESTAMP WITH TIME ZONE,

    FOREIGN KEY (chat_id) REFERENCES chats(id) ON DELETE CASCADE,
    FOREIGN KEY (first_account_id) REFERENCES accounts(id) ON DELETE CASCADE,
    FOREIGN KEY (second_account_id) REFERENCES accounts(id) ON DELETE CASCADE
    );

-- Create function to insert chat before inserting bot
CREATE OR REPLACE FUNCTION create_chat_for_pv()
    RETURNS TRIGGER AS $$
DECLARE
    new_chat_id BIGINT;
BEGIN
    -- Step 1: Insert a new chat record for the bot
    INSERT INTO chats (chat_type, created_at, updated_at)
    VALUES ('pv', NOW(), NOW())
    RETURNING id INTO new_chat_id;

-- Step 2: Set the newly created chat_id to the bot
    NEW.chat_id := new_chat_id;

    -- Return the new bot record with the chat_id set
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger to execute the function after inserting a new bot
CREATE TRIGGER create_chat_for_new_pv
    BEFORE INSERT ON private_chats
    FOR EACH ROW
EXECUTE FUNCTION create_chat_for_pv();