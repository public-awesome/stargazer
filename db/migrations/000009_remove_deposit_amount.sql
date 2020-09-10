-- +migrate Up
ALTER TABLE posts
  DROP COLUMN deposit_amount,
  DROP COLUMN deposit_denom;

ALTER TABLE upvotes
  DROP COLUMN deposit_amount,
  DROP COLUMN deposit_denom;

-- +migrate Down
ALTER TABLE posts
  ADD COLUMN deposit_amount bigint NOT NULL DEFAULT 0,
  ADD COLUMN deposit_denom text NOT NULL DEFAULT '';

ALTER TABLE upvotes
  ADD COLUMN deposit_amount bigint NOT NULL DEFAULT 0,
  ADD COLUMN deposit_denom text NOT NULL DEFAULT '';

