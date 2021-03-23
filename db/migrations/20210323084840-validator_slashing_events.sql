-- +migrate Up
CREATE TABLE slashing_events (
  id bigserial PRIMARY KEY,
  height bigint NOT NULL,
  validator_address text NOT NULL,
  event_type text NOT NULL,
  counter bigint NOT NULL DEFAULT 0,
  reason text NOT NULL DEFAULT '',
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp,
  CONSTRAINT slashing_events_validators_validator_address_fkey FOREIGN KEY (validator_address) REFERENCES validators (address)
);

CREATE INDEX slashing_events_height_address ON slashing_events (height, validator_address);

-- +migrate Down
DROP TABLE slashing_events;

