CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY NOT NULL,
    email VARCHAR,
    phone VARCHAR,
    name VARCHAR,
    password VARCHAR NOT NULL,
    image_url VARCHAR DEFAULT '',
    friend_count INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL
);
