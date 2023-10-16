-- +goose Up
CREATE TABLE IF NOT EXISTS status (
    status_id SERIAL PRIMARY KEY,
    status_name VARCHAR(10) UNIQUE NOT NULL
);

INSERT INTO status (status_name) VALUES
    ('NEW'),
    ('PROCESSING'),
    ('INVALID'),
    ('PROCESSED');

-- +goose Down
DROP TABLE IF EXISTS status;