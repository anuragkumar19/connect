-- +goose Up 
CREATE DOMAIN "ulid" AS bytea;

-- +goose Down
DROP DOMAIN "ulid";