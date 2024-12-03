-- name: IsUsernameAvailable :one
SELECT
    NOT EXISTS (
        SELECT
            1
        FROM
            "users"
        WHERE
            "users"."username" = $1
    )
    AND NOT EXISTS (
        SELECT
            1
        FROM
            "reserved_usernames"
        WHERE
            "reserved_usernames"."username" = $1
    ) AS "is_username_available";