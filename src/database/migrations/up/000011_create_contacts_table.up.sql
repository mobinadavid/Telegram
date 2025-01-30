create table if not exists contacts
(
    nickname varchar(20),
    is_favorite     boolean default false,
    is_blocked      boolean default false,
    account_id       bigint not null
    constraint fk_contacts_account
    references accounts,
    contact_id bigint not null
    constraint fk_contacts_contact
    references accounts,
    primary key (account_id, contact_id)
    );

