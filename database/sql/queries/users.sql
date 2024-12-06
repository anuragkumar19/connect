-- name: IsUsernameAvailable :one
SELECT
    NOT EXISTS (
        SELECT
            1
        FROM
            "users"
        WHERE
            "users"."username" = $1
            AND "users"."is_registered"
    )
    AND NOT EXISTS (
        SELECT
            1
        FROM
            "reserved_usernames"
        WHERE
            "reserved_usernames"."username" = $1
    ) AS "is_username_available";

-- name: CreateUser :exec
INSERT INTO
    "users" (
        "id",
        "is_registered",
        "name",
        "username",
        "primary_email_id",
        "primary_phone_number_id"
    )
VALUES
    ($1, $2, $3, $4, $5, $6);