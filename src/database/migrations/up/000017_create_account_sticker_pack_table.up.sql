create table if not exists account_sticker_packs
(
    used_at         TIMESTAMP DEFAULT NULL ,
    usage_count     bigint,

    account_id      bigint not null
    constraint fk_account_sticker_packs_account
    references accounts,
    sticker_pack_id bigint not null
    constraint fk_account_sticker_packs_sticker_pack
    references sticker_packs,
    primary key (account_id, sticker_pack_id)
    );

