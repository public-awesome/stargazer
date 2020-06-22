
-- +migrate Up
CREATE TABLE posts (
  id BIGSERIAL PRIMARY KEY,
  vendor_id INTEGER NOT NULL,
  author_address VARCHAR(65),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);
-- +migrate Down
DROP TABLE posts;
