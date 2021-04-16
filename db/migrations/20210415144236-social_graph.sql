-- +migrate Up
CREATE TABLE social_graph (
  id bigserial PRIMARY KEY,
  amount bigint NOT NULL,
  buyer_address text NOT NULL,
  creator_address text NOT NULL,
  height bigint NOT NULL,
  username text NOT NULL,
  validator_address text NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp
);

-- +migrate Down
DROP TABLE social_graph;
