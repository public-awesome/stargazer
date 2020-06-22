
-- +migrate Up
CREATE TABLE pre_commits
(
  id SERIAL PRIMARY KEY,
  height INTEGER NOT NULL,
  round INTEGER NOT NULL,
  validator_address TEXT NOT NULL,
  pre_commit_timestamp TIMESTAMP NOT NULL,
  voting_power INTEGER NOT NULL,
  proposer_priority INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP,
  CONSTRAINT pre_commits_validators_validator_address_fkey FOREIGN KEY (validator_address) REFERENCES validators(address),
  CONSTRAINT pre_commits_uq UNIQUE (height, round)
);

-- +migrate Down
DROP TABLE pre_commits;
