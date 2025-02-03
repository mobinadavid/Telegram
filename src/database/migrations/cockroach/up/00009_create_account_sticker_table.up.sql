CREATE TABLE IF NOT EXISTS account_stickers (
     used_at TIMESTAMPTZ DEFAULT NULL,
     usage_count BIGINT,
     account_id UUID NOT NULL,
     sticker_id UUID NOT NULL,
    PRIMARY KEY (account_id, sticker_id),
    CONSTRAINT fk_account_stickers_account FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_account_stickers_sticker FOREIGN KEY (sticker_id) REFERENCES stickers(id) ON DELETE CASCADE
    );