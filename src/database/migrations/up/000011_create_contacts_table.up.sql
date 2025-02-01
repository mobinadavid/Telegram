CREATE TABLE IF NOT EXISTS contacts (
nickname       VARCHAR(20),
is_favorite    BOOLEAN DEFAULT FALSE,
is_blocked     BOOLEAN DEFAULT FALSE,
account_id     BIGINT NOT NULL,
contact_id     BIGINT NOT NULL,
CONSTRAINT fk_contacts_account
FOREIGN KEY (account_id) REFERENCES accounts(id),
CONSTRAINT fk_contacts_contact
FOREIGN KEY (contact_id) REFERENCES accounts(id),
PRIMARY KEY (account_id, contact_id)
);


CREATE OR REPLACE FUNCTION create_chat_on_contact_add()
    RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO chats (account_id, type_id, chat_type)
    VALUES (NEW.account_id, NEW.contact_id, 'pv');

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to create chat after adding contact
CREATE TRIGGER trigger_create_chat_on_contact_add
    AFTER INSERT ON contacts
    FOR EACH ROW
EXECUTE FUNCTION create_chat_on_contact_add();
