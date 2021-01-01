-- +migrate Up 
ALTER TABLE stakes
  DROP CONSTRAINT stake_uq,
  ADD CONSTRAINT stake_uq UNIQUE (vendor_id, post_id, delegator, validator);

-- +migrate Down
ALTER TABLE stakes
  DROP CONSTRAINT stake_uq,
  ADD CONSTRAINT stake_uq UNIQUE (vendor_id, post_id);
