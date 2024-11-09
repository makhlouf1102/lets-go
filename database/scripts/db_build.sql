CREATE TABLE user (
    -- User ID is a UUID
    user_id CHAR(36) PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    first_name TEXT,
    last_name TEXT
);
