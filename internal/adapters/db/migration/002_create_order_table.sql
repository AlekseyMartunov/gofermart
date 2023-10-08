-- +goose Up
CREATE TABLE IF NOT EXISTS "order" (
    order_id SERIAL PRIMARY KEY,
    order_number VARCHAR(50)
);

-- +goose Down
DROP TABLE IF EXISTS "order";