create table if not exists groups
(
    id               bigserial primary key,
    name             VARCHAR(255) NOT NULL,
    description      VARCHAR(255),
    owner_id         bigint NOT NULL,
    members_count    bigint default 1,
    invite_link      text,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone,
    FOREIGN KEY (owner_id) REFERENCES accounts (id) ON DELETE CASCADE

    );


