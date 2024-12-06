-- name: CreatePassword :exec
INSERT INTO
    "passwords" ("id", "is_active", "user_id", "value")
VALUES
    ($1, $2, $3, $4);