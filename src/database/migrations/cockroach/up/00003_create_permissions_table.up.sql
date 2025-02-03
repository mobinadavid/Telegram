CREATE TABLE IF NOT EXISTS permissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name STRING,
    title JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp(),
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
                             );

-- Create index for deleted_at
CREATE INDEX IF NOT EXISTS idx_permissions_deleted_at
    ON permissions (deleted_at);

-- Create unique index for name
CREATE UNIQUE INDEX IF NOT EXISTS idx_permissions_name
    ON permissions (name);