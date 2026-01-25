-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(32) NOT NULL,
    password_hash BYTEA NOT NULL,
    role VARCHAR(16) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT uk_email UNIQUE (email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
