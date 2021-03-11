-- +migrate Up
ALTER TABLE posts
  ADD COLUMN chain_id text NOT NULL default '',
  ADD COLUMN owner text NOT NULL default '',
  ADD COLUMN contract_address text NOT NULL default '',
  ADD COLUMN metadata text NOT NULL default '',
  ADD COLUMN locked boolean NOT NULL DEFAULT FALSE,
  ADD COLUMN parent_id text NOT NULL default '';

-- +migrate Down
ALTER TABLE posts
  DROP COLUMN chain_id,
  DROP COLUMN owner,
  DROP COLUMN contract_address,
  DROP COLUMN metadata,
  DROP COLUMN locked,
  DROP COLUMN parent_id;
