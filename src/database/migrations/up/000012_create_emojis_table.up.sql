create table if not exists emojis
(
    id                    bigserial primary key,
    emoji_url             text not null,
    created_at            timestamp with time zone,
    updated_at            timestamp with time zone,
    deleted_at            timestamp with time zone
    );