-- +goose Up
CREATE TABLE IF NOT EXISTS "order" (
    order_id SERIAL PRIMARY KEY,
    order_number VARCHAR(50) UNIQUE,
    fk_order_status INTEGER REFERENCES status(status_id) DEFAULT 1,
    created_time TIMESTAMP NOT NULL,
    accrual INTEGER,
    fk_user_id INTEGER REFERENCES client(client_id)
);

-- +goose Down
DROP TABLE IF EXISTS "order";