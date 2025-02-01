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