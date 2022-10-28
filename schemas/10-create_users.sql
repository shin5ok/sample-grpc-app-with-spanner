CREATE TABLE users (
  user_id STRING(36) NOT NULL,
  name STRING(MAX) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
) PRIMARY KEY(user_id)
