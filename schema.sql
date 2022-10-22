CREATE TABLE users (
  user_id STRING(36) NOT NULL,
  name STRING(MAX) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
) PRIMARY KEY(user_id);

CREATE TABLE user_items (
  user_id STRING(36) NOT NULL,
  item_id STRING(36) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
) PRIMARY KEY(user_id, item_id),
  INTERLEAVE IN PARENT users ON DELETE CASCADE;

CREATE TABLE items (
  item_id STRING(36) NOT NULL,
  item_name STRING(64) NOT NULL,
  price INT64 NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
) PRIMARY KEY(item_id);
