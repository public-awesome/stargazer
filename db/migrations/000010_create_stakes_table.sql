-- +migrate Up
CREATE TABLE stakes (
  id serial PRIMARY KEY,
  height bigint NOT NULL,
  vendor_id integer NOT NULL,
  post_id text NOT NULL,
  delegator text NOT NULL,
  validator text NOT NULL,
  amount bigint NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp,
  UNIQUE (vendor_id, post_id)
);

-- +migrate Down
DROP TABLE stakes;
