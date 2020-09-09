-- +migrate Up
ALTER TABLE posts
  ADD COLUMN total_voter_count int NOT NULL DEFAULT 0;

-- +migrate Down
ALTER TABLE posts
  DROP COLUMN total_voter_count;