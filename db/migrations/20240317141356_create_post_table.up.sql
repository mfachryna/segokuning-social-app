CREATE TABLE posts (
    id UUID PRIMARY KEY NOT NULL,
    content VARCHAR NOT NULL,
    tags VARCHAR[] NOT NULL,
    user_id UUID REFERENCES users(id) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMPNOT NULL
);