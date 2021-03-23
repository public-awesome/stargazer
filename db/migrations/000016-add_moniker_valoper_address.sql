-- +migrate Up
ALTER TABLE validators
  ADD COLUMN operator_address text NOT NULL DEFAULT '',
  ADD COLUMN moniker text NOT NULL DEFAULT '';

-- +migrate Down
ALTER TABLE validators
  DROP COLUMN operator_address,
  DROP COLUMN moniker;

