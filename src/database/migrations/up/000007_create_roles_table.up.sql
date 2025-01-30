create table if not exists roles
(
    id         bigserial
    primary key,
    title      varchar(255) not null,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
                             );

create index if not exists idx_roles_deleted_at
    on roles (deleted_at);

create unique index if not exists idx_roles_title
    on roles (title);
