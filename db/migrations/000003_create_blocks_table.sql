
-- +migrate Up
CREATE TABLE blocks
(
  height INTEGER PRIMARY KEY,
  hash TEXT NOT NULL UNIQUE,
  num_txs INTEGER NOT NULL DEFAULT 0,
  total_gas INTEGER  NOT NULL DEFAULT 0,
  proposer_address TEXT NOT NULL,
  signatures INTEGER NOT NULL,
  block_timestamp TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP,
  CONSTRAINT blocks_validators_proposer_address_fkey FOREIGN KEY (proposer_address) REFERENCES validators(address)
);

-- +migrate Down
DROP TABLE blocks;
