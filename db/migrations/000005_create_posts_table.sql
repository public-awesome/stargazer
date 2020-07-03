
-- +migrate Up
CREATE TABLE posts (
  id BIGSERIAL PRIMARY KEY,
  vendor_id INTEGER NOT NULL,
  post_id BIGINT NOT NULL,
  creator VARCHAR(65) NOT NULL,
  reward_address VARCHAR(65) NOT NULL,
  deposit TEXT NOT NULL,
  timestamp TIMESTAMP NOT NULL,
  curation_end_time TIMESTAMP NOT NULL,
  body TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);
-- +migrate Down
DROP TABLE posts;
