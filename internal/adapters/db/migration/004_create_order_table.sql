-- +goose Up
CREATE TABLE IF NOT EXISTS orders (
    order_id SERIAL PRIMARY KEY,
    order_number VARCHAR(50) UNIQUE,
    fk_order_status INTEGER REFERENCES status(status_id) DEFAULT 1,
    created_time TIMESTAMP NOT NULL,
    accrual NUMERIC(10, 5) CHECK(accrual >= 0) DEFAULT 0,
    discount NUMERIC(10, 5) CHECK(discount >= 0) DEFAULT 0,
    fk_user_id INTEGER REFERENCES client(client_id)
);

-- +goose Down
DROP TABLE IF EXISTS orders;