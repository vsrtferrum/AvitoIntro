-- +goose Up
-- +goose StatementBegin

CREATE TABLE sales (
    customer_id BIGINT REFERENCES users(id),
    item_id BIGINT REFERENCES shop(id),
    cost BIGINT NOT NULL CHECK (cost > 0),
    date_of_purchase TIMESTAMP NOT NULL
);

CREATE OR REPLACE FUNCTION set_current_date()
RETURNS TRIGGER AS $$
BEGIN
    NEW.date_of_purchase = CURRENT_DATE; 
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_sales_date_of_purchase
BEFORE INSERT ON sales
FOR EACH ROW
EXECUTE FUNCTION set_current_date();
CREATE INDEX sales_idx_customer_id on sales(customer_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS sales_idx_customer_id;
DROP TRIGGER IF EXISTS set_sales_date_of_purchase ON sales;
DROP TABLE sales;

-- +goose StatementEnd