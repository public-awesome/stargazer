-- +migrate Up
ALTER TABLE posts
  ADD COLUMN total_staked_amount bigint NOT NULL DEFAULT 0;

-- +migrate Down
ALTER TABLE posts
  DROP COLUMN total_staked_amount;
