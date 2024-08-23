CREATE TABLE sessions (
    "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "expires_at" TIMESTAMP,
    "name" VARCHAR(255) NOT NULL,
    "message" VARCHAR(255) NOT NULL
)