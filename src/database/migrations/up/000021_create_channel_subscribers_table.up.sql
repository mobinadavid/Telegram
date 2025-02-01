CREATE TABLE IF NOT EXISTS channel_subscribers
(
    channel_id       bigint,
    account_id       bigint,
    is_admin         boolean default false,
    joined_at        timestamp with time zone,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone,

    PRIMARY KEY (channel_id, account_id),
    FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE,
    FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE

    );


-- Function to increment member count
CREATE OR REPLACE FUNCTION increment_subscribers_count_channel()
    RETURNS TRIGGER AS $$
BEGIN
    -- Increment the member count for the group by 1
    UPDATE channels
    SET subscribers_count = subscribers_count + 1
    WHERE id = NEW.channel_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to increment member count after an insert
CREATE TRIGGER increment_subscriber_count_trigger_channel
    AFTER INSERT ON channel_subscribers
    FOR EACH ROW
EXECUTE FUNCTION increment_subscribers_count_channel();


-- Function to decrement member count
CREATE OR REPLACE FUNCTION decrement_subscribers_count_channel()
    RETURNS TRIGGER AS $$
BEGIN
    -- Decrement the member count for the group by 1
    UPDATE channels
    SET subscribers_count = subscribers_count - 1
    WHERE id = OLD.channel_id;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

-- Trigger to decrement member count after a delete
CREATE TRIGGER decrement_subscribers_count_trigger_channel
    AFTER DELETE ON channel_subscribers
    FOR EACH ROW
EXECUTE FUNCTION decrement_subscribers_count_channel();
--------------

CREATE OR REPLACE FUNCTION create_chat_on_channel_add()
    RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO chats (account_id, type_id, chat_type)
    VALUES (
               NEW.account_id,  -- The account that joined the bot
               NEW.channel_id,      -- The bot the account joined
               'channel'          -- The chat type is 'bot'
           );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to create chat after adding contact
CREATE TRIGGER trigger_create_chat_on_channel_add
    AFTER INSERT ON channel_subscribers
    FOR EACH ROW
EXECUTE FUNCTION create_chat_on_channel_add();