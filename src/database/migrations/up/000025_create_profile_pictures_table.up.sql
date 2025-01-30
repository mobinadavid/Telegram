CREATE TABLE IF NOT EXISTS profile_pictures (
  id BIGSERIAL PRIMARY KEY,
  picture_url TEXT NOT NULL,
  owner_id BIGINT NOT NULL,
  owner_type TEXT NOT NULL CHECK (owner_type IN ('account', 'group', 'channel')),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
deleted_at TIMESTAMP WITH TIME ZONE,
   UNIQUE (owner_id, owner_type)
);

-- Trigger function to check if owner_id exists in the correct table
CREATE OR REPLACE FUNCTION enforce_profile_picture_owner() RETURNS TRIGGER AS $$
BEGIN
    CASE NEW.owner_type
        WHEN 'account' THEN
            IF NOT EXISTS (SELECT 1 FROM accounts WHERE id = NEW.owner_id) THEN
                RAISE EXCEPTION 'Invalid owner_id: No matching account';
            END IF;
        WHEN 'group' THEN
            IF NOT EXISTS (SELECT 1 FROM groups WHERE id = NEW.owner_id) THEN
                RAISE EXCEPTION 'Invalid owner_id: No matching group';
            END IF;
        WHEN 'channel' THEN
            IF NOT EXISTS (SELECT 1 FROM channels WHERE id = NEW.owner_id) THEN
                RAISE EXCEPTION 'Invalid owner_id: No matching channel';
            END IF;
        ELSE
            RAISE EXCEPTION 'Invalid owner_type';
        END CASE;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Attach trigger to profile_pictures
CREATE TRIGGER check_profile_picture_owner
    BEFORE INSERT OR UPDATE ON profile_pictures
    FOR EACH ROW EXECUTE FUNCTION enforce_profile_picture_owner();
