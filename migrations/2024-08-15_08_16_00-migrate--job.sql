CREATE TABLE jobs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    name TEXT NOT NULL,
    topic TEXT NULL,
    is_visible TINYINT(1) NOT NULL DEFAULT 1,
    message TEXT NOT NULL
)