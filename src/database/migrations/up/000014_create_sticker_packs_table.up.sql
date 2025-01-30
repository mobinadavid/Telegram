create table if not exists sticker_packs
(
    id                    bigserial primary key,
    name                  varchar(255),
    created_at            timestamp with time zone,
    updated_at            timestamp with time zone,
    deleted_at            timestamp with time zone
    );
