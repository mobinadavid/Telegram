create table if not exists profile
(
    id                       bigserial primary key,
    owner_id                 bigint UNIQUE,
    first_name               varchar(255) default null,
    last_name                varchar(255) default null,
    bio                      varchar(255) default null,
    created_at               timestamp with time zone,
    updated_at               timestamp with time zone,
    deleted_at               timestamp with time zone,
      FOREIGN KEY (owner_id) REFERENCES accounts (id) ON DELETE CASCADE

    );

create index if not exists idx_profile_deleted_at
    on profile (deleted_at);

