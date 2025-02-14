-- +goose Up
-- +goose StatementBegin

CREATE TABLE transfers (
    sender_id BIGINT REFERENCES users(id),
    recipient_id BIGINT REFERENCES users(id),
    cost BIGINT NOT NULL CHECK (cost > 0),
    date_of_purchase TIMESTAMP NOT NULL
);

CREATE OR REPLACE FUNCTION set_current_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.date_of_purchase = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_transfers_date_of_purchase
BEFORE INSERT ON transfers
FOR EACH ROW
EXECUTE FUNCTION set_current_timestamp();
CREATE INDEX transfers_idx_sender_id ON transfers (sender_id);
CREATE INDEX transfers_idx_recipient_id ON transfers(recipient_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS transfers_idx_sender_id ;
DROP TRIGGER IF EXISTS transfers_idx_recipient_id ON transfers;
DROP TABLE transfers;
-- +goose StatementEnd