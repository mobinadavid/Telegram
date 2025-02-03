CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username STRING DEFAULT NULL,
    password STRING NOT NULL,
    mobile STRING NOT NULL UNIQUE,
    email STRING DEFAULT NULL UNIQUE,
    birth_date STRING DEFAULT NULL,
    is_premium BOOL DEFAULT FALSE,
    is_active BOOL DEFAULT TRUE,
    last_seen TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),
    updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),
    deleted_at TIMESTAMP DEFAULT NULL
    );
