-- name: IsPhoneNumberAvailable :one
SELECT
    NOT EXISTS (
        SELECT
            1
        FROM
            "phone_numbers"
        WHERE
            "value" = $1
            AND "is_verified"
    );

-- name: CreatePhoneNumber :exec
INSERT INTO
    "phone_numbers" (
        "id",
        "user_id",
        "value",
        "is_primary",
        "is_verified",
        "last_verified_at"
    )
VALUES
    ($1, $2, $3, $4, $5, $6);