CREATE TABLE reg_confirmations (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    uuid VARCHAR(36),
    user_id INTEGER
)
