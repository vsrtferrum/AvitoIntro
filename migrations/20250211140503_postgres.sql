-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY, 
    name  TEXT NOT NUlL,
    password TEXT NOT NULL, 
    balance  BIGINT CHECK (balance > 0)
);
CREATE INDEX idx_name ON users (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS  idx_name;
DROP TABLE users;
-- +goose StatementEnd
