-- +migrate Up
CREATE TABLE transactions (
  hash TEXT NOT NULL PRIMARY KEY,
  gas_wanted integer NOT NULL DEFAULT 0,
  gas_used integer NOT NULL DEFAULT 0,
  height bigint NOT NULL,
  events jsonb NOT NULL DEFAULT '[]' ::jsonb,
  messages jsonb NOT NULL DEFAULT '[]' ::jsonb,
  fee jsonb NOT NULL DEFAULT '{}' ::jsonb,
  signatures jsonb NOT NULL DEFAULT '[]' ::jsonb,
  memo varchar(256) NOT NULL DEFAULT '',
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp,
  CONSTRAINT transactions_blocks_height_fkey FOREIGN KEY (height) REFERENCES blocks (height)
);

-- +migrate Down
DROP TABLE transactions;

