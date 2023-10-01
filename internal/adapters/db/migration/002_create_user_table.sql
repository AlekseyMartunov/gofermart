-- +goose Up
CREATE TABLE IF NOT EXISTS client (
    client_id SERIAL PRIMARY KEY,
    login VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(50) NOT NUll,
    balance INTEGER DEFAULT 0,
    fk_order_id INTEGER REFERENCES "order"(order_id)
);
-- +goose Down
DROP TABLE IF EXISTS client;