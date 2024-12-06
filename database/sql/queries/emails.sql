-- name: IsEmailAvailable :one
SELECT
    NOT EXISTS (
        SELECT
            1
        FROM
            "emails"
        WHERE
            "value" = $1
            AND "is_verified"
    );

-- name: CreateEmail :exec
INSERT INTO
    "emails" (
        "id",
        "user_id",
        "value",
        "is_primary",
        "is_verified",
        "last_verified_at"
    )
VALUES
    ($1, $2, $3, $4, $5, $6);