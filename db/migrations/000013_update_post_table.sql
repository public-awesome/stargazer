-- +migrate Up
ALTER TABLE posts
  ADD COLUMN total_staked_amount bigint NOT NULL DEFAULT 0,
  ADD COLUMN total_staked_denom text NOT NULL DEFAULT 'ustarx';

-- +migrate Down
ALTER TABLE posts
  DROP COLUMN total_staked_amount,
  DROP COLUMN total_staked_denom;
