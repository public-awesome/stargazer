-- +migrate Up
ALTER TABLE posts
  ADD COLUMN total_voter_count int DEFAULT 0;

-- +migrate Down
ALTER TABLE posts
  DROP COLUMN total_voter_count;