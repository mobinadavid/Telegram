create table if not exists setting
(
    id                       bigserial primary key,
    owner_id                 bigint UNIQUE,
    language                 TEXT NOT NULL DEFAULT 'en' CHECK (language IN ('fa', 'en')),
    created_at               timestamp with time zone,
    updated_at               timestamp with time zone,
    deleted_at               timestamp with time zone,
      FOREIGN KEY (owner_id) REFERENCES accounts (id) ON DELETE CASCADE

    );

create index if not exists idx_setting_deleted_at
    on setting (deleted_at);

