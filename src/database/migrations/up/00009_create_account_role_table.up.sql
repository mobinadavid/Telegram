create table if not exists account_role
(
    account_id bigint not null
    constraint fk_role_account_account
    references accounts,
    role_id bigint not null
    constraint fk_role_account_role
    references roles,
    primary key (account_id, role_id)
    );

