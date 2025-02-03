CREATE TABLE IF NOT EXISTS group_members (
     group_id UUID,
     account_id UUID,
     is_admin BOOLEAN DEFAULT FALSE,
     joined_at TIMESTAMPTZ DEFAULT current_timestamp(),
    updated_at TIMESTAMPTZ DEFAULT current_timestamp(),
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (group_id, account_id),
    CONSTRAINT fk_group_members_group FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
    CONSTRAINT fk_group_members_account FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE
    );