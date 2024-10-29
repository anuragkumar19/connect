-- name: GetRegistrationFlow :one
SELECT
    *
FROM
    "registration_flows"
WHERE
    "id" = $1;

-- name: CreateRegistrationFlow :one
INSERT INTO
    "registration_flows" ("created_at", "updated_at", "expire_at", "csrf_token", "state")
VALUES
    ($1, $2, $3, $4, $5)
RETURNING
    *;