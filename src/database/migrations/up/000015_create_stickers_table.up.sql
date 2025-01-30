create table if not exists stickers
(
    id                    bigserial primary key,
    sticker_url           text not null,
    sticker_pack_id       bigint,
    emoji_id              bigint,
    created_at            timestamp with time zone,
    updated_at            timestamp with time zone,
    deleted_at            timestamp with time zone,
    FOREIGN KEY (sticker_pack_id) REFERENCES sticker_packs (id) ON DELETE CASCADE,
    FOREIGN KEY (emoji_id) REFERENCES emojis  (id) ON DELETE CASCADE

    );

create index if not exists idx_stickers_deleted_at
    on stickers (deleted_at);
