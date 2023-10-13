-- +goose Up
CREATE TABLE IF NOT EXISTS "order" (
    order_id SERIAL PRIMARY KEY,
    order_number VARCHAR(50) UNIQUE,
    order_status VARCHAR(50),
    created_time TIME NOT NULL,
    fk_user_id INTEGER REFERENCES client(client_id)
);

-- +goose Down
DROP TABLE IF EXISTS "order";