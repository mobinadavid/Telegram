CREATE TYPE content_type_enum AS ENUM ('image', 'video', 'audio', 'text', 'link');

CREATE TABLE IF NOT EXISTS stories (
   id                BIGSERIAL PRIMARY KEY,
   owner_id          BIGINT,
   content_type      content_type_enum NOT NULL,
   content_url       TEXT NOT NULL,
   caption           TEXT,
   views_count      bigint default 0,
    expires_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW() + INTERVAL '1 day',
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at  TIMESTAMP WITH TIME ZONE,

FOREIGN KEY (owner_id) REFERENCES accounts(id) ON DELETE CASCADE

    );
CREATE INDEX IF NOT EXISTS idx_stories_owner_id ON stories(owner_id);

CREATE OR REPLACE FUNCTION delete_expired_stories()
    RETURNS TRIGGER AS $$
BEGIN
    IF NEW.expires_at < NOW() THEN
        DELETE FROM stories WHERE id = NEW.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER delete_expired_story_trigger
    AFTER INSERT OR UPDATE ON stories
    FOR EACH ROW
EXECUTE FUNCTION delete_expired_stories();
