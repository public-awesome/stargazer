-- +migrate Up
CREATE TABLE blocks (
  height bigint PRIMARY KEY,
  hash TEXT NOT NULL UNIQUE,
  num_txs integer NOT NULL DEFAULT 0,
  total_gas bigint NOT NULL DEFAULT 0,
  proposer_address text NOT NULL,
  signatures integer NOT NULL,
  block_timestamp timestamp NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp,
  CONSTRAINT blocks_validators_proposer_address_fkey FOREIGN KEY (proposer_address) REFERENCES validators (address)
);

-- +migrate Down
DROP TABLE blocks;

