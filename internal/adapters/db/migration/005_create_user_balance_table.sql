-- +goose Up
CREATE TABLE IF NOT EXISTS user_balance (
    id SERIAL PRIMARY KEY,
    fk_user_id INTEGER REFERENCES client(client_id),
    bonuses NUMERIC(10, 5) CHECK(bonuses > 0),
    withdrawn  NUMERIC(10, 5) CHECK(withdrawn > 0)
);

-- +goose Down
DROP TABLE IF EXISTS user_balance;