CREATE TABLE IF NOT EXISTS bot_subscribers
(
    bot_id           bigint,
    account_id       bigint,
    is_admin         boolean default false,
    joined_at        timestamp with time zone,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone,

    PRIMARY KEY (bot_id, account_id),
    FOREIGN KEY (bot_id) REFERENCES bots(id) ON DELETE CASCADE,
    FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE

    );

-- Function to increment member count
CREATE OR REPLACE FUNCTION increment_subscribers_count()
RETURNS TRIGGER AS $$
BEGIN
    -- Increment the member count for the group by 1
UPDATE bots
SET subscribers_count = subscribers_count + 1
WHERE id = NEW.bot_id;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to increment member count after an insert
CREATE TRIGGER increment_subscriber_count_trigger
    AFTER INSERT ON bot_subscribers
    FOR EACH ROW
    EXECUTE FUNCTION increment_subscribers_count();

-- Function to decrement member count
CREATE OR REPLACE FUNCTION decrement_subscribers_count()
RETURNS TRIGGER AS $$
BEGIN
    -- Decrement the member count for the group by 1
UPDATE bots
SET subscribers_count = subscribers_count - 1
WHERE id = OLD.bot_id;
RETURN OLD;
END;
$$ LANGUAGE plpgsql;

-- Trigger to decrement member count after a delete
CREATE TRIGGER decrement_subscribers_count_trigger
    AFTER DELETE ON bot_subscribers
    FOR EACH ROW
    EXECUTE FUNCTION decrement_subscribers_count();
--------------
CREATE OR REPLACE FUNCTION create_bot_chat_on_subscribe()
    RETURNS TRIGGER AS $$
BEGIN
    -- Insert into chats table only when a new bot subscriber joins
    INSERT INTO chats (account_id, chat_id, chat_type)
    VALUES (
               NEW.account_id,
               NEW.bot_id,
               'bot'
           );

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;



CREATE TRIGGER trigger_create_chat_on_channel_join
    AFTER INSERT ON bot_subscribers
    FOR EACH ROW
EXECUTE FUNCTION create_bot_chat_on_subscribe();