CREATE TABLE IF NOT EXISTS roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title STRING NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp(),
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
                             );

-- Create index for deleted_at
CREATE INDEX IF NOT EXISTS idx_roles_deleted_at
    ON roles (deleted_at);

-- Create unique index for title
CREATE UNIQUE INDEX IF NOT EXISTS idx_roles_title
    ON roles (title);
