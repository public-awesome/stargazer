-- +migrate Up
ALTER TABLE posts
  ADD COLUMN total_upvote_reward_amount bigint NOT NULL DEFAULT 0,
  ADD COLUMN total_upvote_reward_denom text NOT NULL DEFAULT 'ustarx';

-- +migrate Down
ALTER TABLE posts
  DROP COLUMN total_upvote_reward_amount,
  DROP COLUMN total_upvote_reward_denom;
