-- +migrate Up
CREATE TABLE block_signatures (
  id serial PRIMARY KEY,
  height bigint NOT NULL,
  round integer NOT NULL,
  validator_address text NOT NULL,
  flag int NOT NULL DEFAULT 0,
  timestamp timestamp NOT NULL,
  hash TEXT NOT NULL,
  voting_power integer NOT NULL,
  proposer_priority integer NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp,
  CONSTRAINT block_signatures_validators_validator_address_fkey FOREIGN KEY (validator_address) REFERENCES validators (address),
  CONSTRAINT block_signatures_uq UNIQUE (height, round, validator_address)
);

-- +migrate Down
DROP TABLE block_signatures;

