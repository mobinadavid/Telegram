create table if not exists contacts
(
    nickname varchar(20),
    is_favorite     boolean default false,
    is_blocked      boolean default false,
    account_id       bigint not null
    constraint fk_contacts_account
    references accounts,
    contact_id bigint not null
    constraint fk_contacts_contact
    references accounts,
    primary key (account_id, contact_id)
    );

CREATE OR REPLACE FUNCTION create_chat_on_join()
    RETURNS TRIGGER AS $$
BEGIN
INSERT INTO chats (account_id, chat_id, chat_type)
VALUES (NEW.account_id, NEW.contact_id, 'account')
    ON CONFLICT DO NOTHING;

RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_create_chat_on_channel_join
    AFTER INSERT ON contacts
    FOR EACH ROW
    EXECUTE FUNCTION create_chat_on_join();