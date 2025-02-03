CREATE TABLE IF NOT EXISTS account_role (
   account_id UUID NOT NULL,
   role_id UUID NOT NULL,
    PRIMARY KEY (account_id, role_id),
    CONSTRAINT fk_role_account_account
    FOREIGN KEY (account_id)
    REFERENCES accounts(id),
    CONSTRAINT fk_role_account_role
    FOREIGN KEY (role_id)
    REFERENCES roles(id)
    );
