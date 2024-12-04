-- name: GetRateLimitBucketForUpdate :one
SELECT
    *
FROM
    "rate_limit_buckets"
WHERE
    "id" = $1
FOR UPDATE;

-- name: CreateRateLimitBucket :one
INSERT INTO
    "rate_limit_buckets" (
        "id",
        "last_reset_at",
        "consumed",
        "last_consumed_at"
    )
VALUES
    ($1, $2, $3, $4)
ON CONFLICT DO NOTHING
RETURNING
    *;

-- name: UpdateRateLimitBucket :exec
UPDATE "rate_limit_buckets"
SET
    "last_reset_at" = $1,
    "consumed" = $2,
    "last_consumed_at" = $3
WHERE
    "id" = $4;