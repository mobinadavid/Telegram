create table if not exists frequently_asked_questions
(
    id          bigserial primary key,
    question    varchar(255) not null,
    answer      text not null,
    is_active   boolean default true,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone
                              );

create index if not exists idx_frequently_asked_questions_deleted_at
    on frequently_asked_questions (deleted_at);
