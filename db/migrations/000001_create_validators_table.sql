
-- +migrate Up
CREATE TABLE validators(
  address TEXT NOT NULL PRIMARY KEY,
  pub_key TEXT NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);

-- +migrate Down
DROP TABLE validators;
