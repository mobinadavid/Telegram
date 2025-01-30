CREATE TYPE notification AS ENUM ('default','disabled', 'only_on_silent');

create table if not exists notification_and_sound
(
    id                       bigserial primary key,
    setting_id               bigint UNIQUE,
    vibrate                  notification not null default 'default',
    ringtone                 notification not null default 'default',
    group_notif              boolean default true,
    channel_notif            boolean default true,
    created_at               timestamp with time zone,
    updated_at               timestamp with time zone,
    deleted_at               timestamp with time zone,
    FOREIGN KEY (setting_id) REFERENCES setting (id) ON DELETE CASCADE

    );


create index if not exists idx_notification_and_sound_deleted_at
    on notification_and_sound (deleted_at);

