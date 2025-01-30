create table if not exists permissions
(
    id    bigserial
        primary key,
    name  varchar(255),
    title jsonb,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);

create index if not exists idx_permissions_deleted_at
    on permissions (deleted_at);

create unique index idx_permissions_name
    on permissions (name);


