CREATE TABLE IF NOT EXISTS profile_pictures (
  id bigserial PRIMARY KEY,
  picture_url text NOT NULL,
  owner_id bigint NOT NULL,
  owner_type text NOT NULL CHECK (owner_type IN ('account', 'group')),
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    CONSTRAINT fk_account FOREIGN KEY (owner_id) REFERENCES accounts(id) ON DELETE CASCADE
 --   CONSTRAINT fk_group FOREIGN KEY (owner_id) REFERENCES groups(id) ON DELETE CASCADE
    );

CREATE OR REPLACE FUNCTION validate_owner_id()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.owner_type = 'account' AND NOT EXISTS (SELECT 1 FROM accounts WHERE id = NEW.owner_id) THEN
        RAISE EXCEPTION 'Invalid owner_id for account';

--     ELSIF NEW.owner_type = 'group' AND NOT EXISTS (SELECT 1 FROM groups WHERE id = NEW.owner_id) THEN
--         RAISE EXCEPTION 'Invalid owner_id for group';
END IF;

RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER check_owner_id_before_insert_update
    BEFORE INSERT OR UPDATE ON profile_pictures
                         FOR EACH ROW EXECUTE FUNCTION validate_owner_id();

create index if not exists idx_profile_pictures_created_at
    on profile_pictures (created_at);

