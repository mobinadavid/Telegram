CREATE TABLE IF NOT EXISTS groups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    owner_id UUID NOT NULL,
    members_count BIGINT DEFAULT 1,
    invite_link TEXT,
    created_at TIMESTAMPTZ DEFAULT current_timestamp(),
    updated_at TIMESTAMPTZ DEFAULT current_timestamp(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT fk_groups_owner FOREIGN KEY (owner_id) REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_groups_chat FOREIGN KEY (id) REFERENCES chats(id) ON DELETE CASCADE
    );