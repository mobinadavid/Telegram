create table if not exists accounts
(
    id                           bigserial    primary key,
    username                     varchar(255) default NULL::character varying,
    password                     text,
    mobile                       varchar(100) not null,
    email                        varchar(100) default NULL::character varying,
    birth_date                   varchar(100) default NULL::character varying,
    is_premium                   boolean      default false,
    is_active                    boolean      default true,
    last_seen TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
    );

create index if not exists idx_accounts_deleted_at
    on accounts (deleted_at);

create unique index if not exists idx_accounts_mobile
    on accounts (mobile);

create unique index if not exists idx_accounts_email
    on accounts (email);

---------------------------------
CREATE OR REPLACE FUNCTION create_profile_on_account_creation()
    RETURNS trigger AS
    $$
BEGIN
        INSERT INTO profile (owner_id, first_name, last_name, bio, created_at, updated_at)
        VALUES (NEW.id, NEW.username, '', '', NOW(), NOW());
        RETURN NEW;
END;
    $$ LANGUAGE plpgsql;

-----------------------------------
CREATE OR REPLACE FUNCTION create_profile_and_settings_on_account_creation()
    RETURNS trigger AS
$$
DECLARE
    new_setting_id bigint; -- Use bigint instead of INT64
BEGIN
    -- Create profile
    INSERT INTO profile (owner_id, first_name, last_name, bio, created_at, updated_at)
    VALUES (NEW.id, NEW.username, '', '', NOW(), NOW());

    -- Create default settings and store the setting_id
    INSERT INTO setting (owner_id, language, created_at, updated_at)
    VALUES (NEW.id, 'en', NOW(), NOW())
    RETURNING id INTO new_setting_id;

    -- Create privacy_and_security settings
    INSERT INTO privacy_and_security (setting_id, last_seen_visibility, profile_visibility,
                                      phone_number_visibility, bio_visibility, two_step_verification, passcode_lock, created_at, updated_at)
    VALUES (new_setting_id, 'everyone', 'everyone', 'everyone', 'everyone', false, false, NOW(), NOW());

    -- Create notification_and_sound settings
    INSERT INTO notification_and_sound (setting_id, vibrate, ringtone, group_notif, channel_notif, created_at, updated_at)
    VALUES (new_setting_id, 'default', 'default', true, true, NOW(), NOW());

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;




CREATE TRIGGER trigger_create_profile
    AFTER INSERT ON accounts
    FOR EACH ROW
    EXECUTE FUNCTION create_profile_and_settings_on_account_creation();

-- CREATE TRIGGER trigger_create_setting
--     AFTER INSERT ON accounts
--     FOR EACH ROW
-- EXECUTE FUNCTION create_default_settings();
