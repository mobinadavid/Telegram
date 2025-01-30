CREATE TYPE visibility AS ENUM ('everyone', 'contacts', 'nobody');

create table if not exists privacy_and_security
(
    id                       bigserial primary key,
    setting_id               bigint UNIQUE,
    last_seen_visibility     visibility not null default 'everyone',
    profile_visibility       visibility not null default 'everyone',
    phone_number_visibility  visibility not null default 'everyone',
    bio_visibility           visibility not null default 'everyone',
    two_step_verification    boolean default false,
    passcode_lock            boolean default false,
    created_at               timestamp with time zone,
    updated_at               timestamp with time zone,
    deleted_at               timestamp with time zone,
       FOREIGN KEY (setting_id) REFERENCES setting (id) ON DELETE CASCADE

    );


create index if not exists idx_privacy_and_security_deleted_at
    on privacy_and_security (deleted_at);

