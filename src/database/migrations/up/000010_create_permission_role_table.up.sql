create table if not exists permission_role
(
    role_id       bigint not null
    constraint fk_permission_role_role
    references roles,
    permission_id bigint not null
    constraint fk_permission_role_permission
    references permissions,
    primary key (role_id, permission_id)
    );

