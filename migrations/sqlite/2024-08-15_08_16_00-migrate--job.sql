CREATE TABLE jobs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    topic VARCHAR(255) NULL,
    is_visible TINYINT(1) NOT NULL DEFAULT 1,
    message BLOB
)