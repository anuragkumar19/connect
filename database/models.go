// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	ulid "github.com/oklog/ulid/v2"
)

type Email struct {
	ID                               ulid.ULID
	CreatedAt                        time.Time
	UpdatedAt                        time.Time
	Version                          int32
	UserID                           ulid.ULID
	Value                            string
	IsPrimary                        bool
	IsVerified                       bool
	LastVerifiedAt                   time.Time
	VerificationCodeSentCount        int32
	VerificationCodeLastSentAt       pgtype.Timestamptz
	VerificationCodeSentCountResetAt pgtype.Timestamptz
}

type Password struct {
	ID                           ulid.ULID
	CreatedAt                    time.Time
	AbandonedAt                  pgtype.Timestamptz
	IsActive                     bool
	UserID                       ulid.ULID
	Value                        []byte
	IncorrectAttemptCount        int32
	LastIncorrectAttemptAt       pgtype.Timestamptz
	IncorrectAttemptCountResetAt pgtype.Timestamptz
}

type PhoneNumber struct {
	ID                               ulid.ULID
	CreatedAt                        time.Time
	UpdatedAt                        time.Time
	Version                          int32
	UserID                           ulid.ULID
	Value                            string
	IsPrimary                        bool
	IsVerified                       bool
	LastVerifiedAt                   time.Time
	VerificationCodeSentCount        int32
	VerificationCodeLastSentAt       pgtype.Timestamptz
	VerificationCodeSentCountResetAt pgtype.Timestamptz
}

type User struct {
	ID                   ulid.ULID
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Version              int32
	Name                 string
	Username             string
	Dob                  pgtype.Date
	Bio                  pgtype.Text
	Website              pgtype.Text
	Location             pgtype.Text
	AvatarPath           pgtype.Text
	BannerPath           pgtype.Text
	LastActiveAt         time.Time
	PrimaryEmailID       ulid.ULID
	PrimaryPhoneNumberID ULIDValue
	IsMfaEnabled         bool
	IsBot                bool
	ProfileLocked        bool
	ShowDob              bool
}
