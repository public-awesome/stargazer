-- +migrate Up
ALTER TABLE posts
  DROP COLUMN total_upvote_reward_denom,
  ADD COLUMN total_upvote_reward_denom text NOT NULL DEFAULT 'ustarx';

-- +migrate Down
ALTER TABLE posts
  DROP COLUMN total_upvote_reward_denom,
  ADD COLUMN total_upvote_reward_denom text NOT NULL;
