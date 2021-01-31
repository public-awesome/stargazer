-- +migrate Up
ALTER TABLE posts
  ADD COLUMN total_upvote_reward int NOT NULL DEFAULT 0;

-- +migrate Down
ALTER TABLE posts
  DROP COLUMN total_upvote_reward;