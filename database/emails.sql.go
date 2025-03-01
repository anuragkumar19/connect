// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: emails.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	ulid "github.com/oklog/ulid/v2"
)

const createEmail = `-- name: CreateEmail :exec
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
    ($1, $2, $3, $4, $5, $6)
`

type CreateEmailParams struct {
	ID             ulid.ULID
	UserID         ulid.ULID
	Value          string
	IsPrimary      bool
	IsVerified     bool
	LastVerifiedAt pgtype.Timestamptz
}

func (q *Queries) CreateEmail(ctx context.Context, arg *CreateEmailParams) error {
	_, err := q.db.Exec(ctx, createEmail,
		arg.ID,
		arg.UserID,
		arg.Value,
		arg.IsPrimary,
		arg.IsVerified,
		arg.LastVerifiedAt,
	)
	return err
}

const isEmailAvailable = `-- name: IsEmailAvailable :one
SELECT
    NOT EXISTS (
        SELECT
            1
        FROM
            "emails"
        WHERE
            "value" = $1
            AND "is_verified"
    )
`

func (q *Queries) IsEmailAvailable(ctx context.Context, value string) (bool, error) {
	row := q.db.QueryRow(ctx, isEmailAvailable, value)
	var not_exists bool
	err := row.Scan(&not_exists)
	return not_exists, err
}
