
-- +migrate Up
CREATE TABLE block_signatures
(
  id SERIAL PRIMARY KEY,
  height BIGINT NOT NULL,
  round INTEGER NOT NULL,
  validator_address TEXT NOT NULL,
  flag INT NOT NULL DEFAULT 0,
  timestamp TIMESTAMP NOT NULL,
  hash TEXT NOT NULL UNIQUE,
  voting_power INTEGER NOT NULL,
  proposer_priority INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP,
  CONSTRAINT block_signatures_validators_validator_address_fkey FOREIGN KEY (validator_address) REFERENCES validators(address),
  CONSTRAINT block_signatures_uq UNIQUE (height, round, validator_address)
);

-- +migrate Down
DROP TABLE block_signatures;
