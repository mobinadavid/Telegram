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
