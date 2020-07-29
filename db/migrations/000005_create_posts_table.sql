-- +migrate Up
CREATE TABLE posts (
  id uuid PRIMARY KEY,
  height bigint NOT NULL,
  vendor_id integer NOT NULL,
  post_id text NOT NULL,
  creator text NOT NULL,
  reward_address text NOT NULL,
  deposit_amount bigint NOT NULL,
  deposit_denom text NOT NULL,
  timestamp timestamp NOT NULL,
  curation_end_time timestamp NOT NULL,
  body text NOT NULL,
  total_votes int NOT NULL DEFAULT 0,
  total_votes_amount bigint NOT NULL DEFAULT 0,
  total_votes_denom text NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp,
  UNIQUE (vendor_id, post_id)
);

-- +migrate Down
DROP TABLE posts;

