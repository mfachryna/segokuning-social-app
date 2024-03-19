CREATE TABLE IF NOT EXISTS friends (
  user_id UUID REFERENCES users(id) NOT NULL,
  friend_id UUID REFERENCES users(id) NOT NULL,
  created_at TIMESTAMP NOT NULL
)
