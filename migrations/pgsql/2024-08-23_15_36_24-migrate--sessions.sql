CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    message BYTEA
)