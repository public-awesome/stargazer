-- +migrate Up
CREATE TABLE upvotes (
  id bigserial PRIMARY KEY,
  vendor_id integer NOT NULL,
  post_id text NOT NULL,
  creator text NOT NULL,
  reward_address text NOT NULL,
  deposit text NOT NULL,
  timestamp timestamp NOT NULL,
  curation_end_time timestamp NOT NULL,
  body text NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp
);

-- +migrate Down
DROP TABLE upvotes;

