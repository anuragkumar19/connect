-- name: GetRateLimitBucket :one
SELECT
    *
FROM
    "rate_limit_buckets"
WHERE
    "id" = $1;

-- name: CreateRateLimitBucket :exec
INSERT INTO
    "rate_limit_buckets" ("id", "last_reset_at", "version", "consumed", "last_consumed_at")
VALUES
    ($1, $2, $3, $4, $5);

-- name: UpdateRateLimitBucket :one
UPDATE "rate_limit_buckets"
SET
    "last_reset_at" = $1,
    "consumed" = $2,
    "last_consumed_at" = $3,
    "version" = "version" + 1
WHERE
    "id" = $4
    AND "version" = $5
RETURNING
    *;