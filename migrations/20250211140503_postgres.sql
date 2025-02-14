-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY, 
    name  TEXT UNIQUE NOT NUlL,
    password TEXT NOT NULL, 
    balance  BIGINT CHECK (balance > 0)
);
CREATE INDEX user_idx_name ON users (name);
CREATE INDEX user_idx_id ON users (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS  user_idx_name;
DROP INDEX IF EXISTS  user_idx_id;
DROP TABLE users;
-- +goose StatementEnd
