-- +goose Up
-- +goose StatementBegin
CREATE TABLE shop (
    id BIGSERIAL PRIMARY KEY, 
    name  text NOT NULL UNIQUE,  
    cost  INTEGER not NULL CHECK (cost > 0));
CREATE INDEX shop_idx_id ON shop (id);
CREATE INDEX shop_idx_name ON shop (name);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS shop_idx_id;
DROP INDEX IF EXISTS shop_idx_name;
DROP TABLE shop;
-- +goose StatementEnd
