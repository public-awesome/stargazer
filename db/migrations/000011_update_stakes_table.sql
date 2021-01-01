-- +migrate Up 
ALTER TABLE stakes
  DROP CONSTRAINT stakes_vendor_id_post_id_key,
  ADD UNIQUE (vendor_id, post_id, delegator, validator);

-- +migrate Down
ALTER TABLE stakes
  DROP CONSTRAINT stakes_vendor_id_post_id_delegator_validator_key,
  ADD UNIQUE (vendor_id, post_id);
