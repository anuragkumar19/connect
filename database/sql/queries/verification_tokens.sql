-- name: CreateVerificationToken :exec
INSERT INTO
    "verification_tokens" (
        "token",
        "id",
        "expire_at",
        "user_id",
        "email_id",
        "phone_number_id",
        "otp"
    )
VALUES
    ($1, $2, $3, $4, $5, $6, $7);