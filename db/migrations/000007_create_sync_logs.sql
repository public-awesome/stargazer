-- +migrate Up
CREATE TABLE sync_logs (
  block_height bigint PRIMARY KEY,
  processed boolean NOT NULL DEFAULT FALSE,
  retries int NOT NULL DEFAULT 0,
  step int NOT NULL DEFAULT 0,
  next_retry timestamp,
  synced_at timestamp,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp
);

-- +migrate Down
DROP TABLE sync_logs;

