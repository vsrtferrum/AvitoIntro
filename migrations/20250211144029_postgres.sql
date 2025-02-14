-- +goose Up
-- +goose StatementBegin
CREATE TABLE shop (
    id BIGSERIAL PRIMARY KEY, 
    name  text NOT NULL,  
    cost  INTEGER not NULL CHECK (cost > 0));
CREATE INDEX idx_id ON shop (id);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_id;
DROP TABLE shop;
-- +goose StatementEnd
