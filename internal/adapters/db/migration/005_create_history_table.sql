-- +goose Up
CREATE TABLE IF NOT EXISTS history (
    history_id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(order_id),
    created_time TIMESTAMP NOT NULL DEFAULT now(),
    amount NUMERIC(10, 5) CHECK(amount >= 0) DEFAULT 0,
    fk_user_id INTEGER REFERENCES client(client_id)
);

-- +goose Down
DROP TABLE IF EXISTS history;