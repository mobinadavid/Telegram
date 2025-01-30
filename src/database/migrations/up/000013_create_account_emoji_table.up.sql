create table if not exists account_emoji
(
    used_at         TIMESTAMP DEFAULT NULL ,
    usage_count     bigint,

    account_id      bigint not null
    constraint fk_account_emoji_account
    references accounts,
    emoji_id bigint not null
    constraint fk_account_emoji_emoji
    references emojis,
    primary key (account_id, emoji_id)
    );

