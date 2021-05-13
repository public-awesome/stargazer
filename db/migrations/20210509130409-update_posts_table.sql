-- +migrate Up
ALTER TABLE posts
  ADD COLUMN tx TEXT NOT NULL DEFAULT '';

-- +migrate Down
ALTER TABLE posts
  DROP COLUMN tx;
