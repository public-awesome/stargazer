-- +migrate Up
CREATE TABLE validators (
  address text NOT NULL PRIMARY KEY,
  pub_key text NOT NULL UNIQUE,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp
);

-- +migrate Down
DROP TABLE validators;

