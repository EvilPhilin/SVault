CREATE TABLE IF NOT EXISTS vaults (
    id CHAR(128) PRIMARY KEY,
    access_token_digest CHAR(128) NOT NULL,
    key CHAR(128) NOT NULL,
    path VARCHAR(150) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);