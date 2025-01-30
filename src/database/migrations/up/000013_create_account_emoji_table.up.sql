create table if not exists account_emoji
(
    used_at         TIMESTAMP DEFAULT NULL ,
    usage_count     bigint,

    account_id      bigint not null
    constraint fk_account_emoji_account
    references accounts on delete cascade,
    emoji_id bigint not null
    constraint fk_account_emoji_emoji
    references emojis  on delete cascade,
    primary key (account_id, emoji_id)
    );

