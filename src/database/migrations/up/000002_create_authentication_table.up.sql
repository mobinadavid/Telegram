create table if not exists authentication
(
    id                       bigserial primary key,
    owner_id                 bigint UNIQUE,
    access_token             text not null,
    access_token_expires_at  timestamp with time zone,
    refresh_token            text not null,
    ip                       varchar(255) not null,
    refresh_token_expires_at timestamp with time zone,
    last_used_at             timestamp with time zone,
    created_at               timestamp with time zone,
    updated_at               timestamp with time zone,
    deleted_at               timestamp with time zone,
  FOREIGN KEY (owner_id) REFERENCES accounts (id) ON DELETE CASCADE

    );

create index if not exists idx_authentication_deleted_at
    on authentication (deleted_at);

