
-- +migrate Up
CREATE TABLE transactions
(
  hash TEXT NOT NULL PRIMARY KEY,
  gas_wanted INTEGER NOT NULL DEFAULT 0,
  gas_used INTEGER NOT NULL DEFAULT 0,
  height BIGINT NOT NULL,
  events jsonb NOT NULL DEFAULT '[]'::jsonb,
  messages jsonb NOT NULL DEFAULT '[]'::jsonb,
  fee jsonb NOT NULL DEFAULT '{}'::jsonb,
  signatures jsonb NOT NULL DEFAULT '[]'::jsonb,
  memo VARCHAR(256) NOT NULL DEFAULT '',
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP,
  CONSTRAINT transactions_blocks_height_fkey FOREIGN KEY (height) REFERENCES blocks(height)
);

-- +migrate Down
DROP TABLE transactions;
