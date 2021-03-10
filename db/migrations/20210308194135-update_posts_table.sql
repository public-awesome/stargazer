-- +migrate Up
ALTER TABLE posts
  ADD COLUMN body_hash text NOT NULL default '';

-- +migrate Down
ALTER TABLE posts
  DROP COLUMN body_hash;
