CREATE TABLE IF NOT EXISTS group_members
(
    group_id         bigint,
    account_id       bigint,
    is_muted         boolean default false,
    is_archived      boolean default false,
    is_admin         boolean default false,
    joined_at        timestamp with time zone,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone,

    PRIMARY KEY (group_id, account_id),
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
    FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE

    );

-- Function to increment member count
CREATE OR REPLACE FUNCTION increment_member_count()
RETURNS TRIGGER AS $$
BEGIN
    -- Increment the member count for the group by 1
UPDATE groups
SET members_count = members_count + 1
WHERE id = NEW.group_id;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to increment member count after an insert
CREATE TRIGGER increment_member_count_trigger
    AFTER INSERT ON group_members
    FOR EACH ROW
    EXECUTE FUNCTION increment_member_count();

-- Function to decrement member count
CREATE OR REPLACE FUNCTION decrement_member_count()
RETURNS TRIGGER AS $$
BEGIN
    -- Decrement the member count for the group by 1
UPDATE groups
SET members_count = members_count - 1
WHERE id = OLD.group_id;
RETURN OLD;
END;
$$ LANGUAGE plpgsql;

-- Trigger to decrement member count after a delete
CREATE TRIGGER decrement_member_count_trigger
    AFTER DELETE ON group_members
    FOR EACH ROW
    EXECUTE FUNCTION decrement_member_count();
