-- +migrate Up
CREATE TABLE upvote_rewards (
  id serial PRIMARY KEY,
  height bigint NOT NULL,
  vendor_id integer NOT NULL,
  post_id text NOT NULL,
  reward_address text NOT NULL,
  reward_amount bigint NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp
);

-- +migrate Down
DROP TABLE upvote_rewards;

