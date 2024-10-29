// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: registration_flows.sql

package database

import (
	"context"
	"time"
)

const createRegistrationFlow = `-- name: CreateRegistrationFlow :one
INSERT INTO
    "registration_flows" ("created_at", "updated_at", "expire_at", "csrf_token", "state")
VALUES
    ($1, $2, $3, $4, $5)
RETURNING
    id, created_at, updated_at, version, expire_at, csrf_token, state, email_value, email_is_valid, email_error_message, email_code, email_code_expiry, email_code_incorrect_verification_attempts, email_code_remaining_verification_attempts, email_code_is_verified, email_code_error_message, phone_number_value, phone_number_is_valid, phone_number_error_message, phone_number_code, phone_number_code_expiry, phone_number_code_incorrect_verification_attempts, phone_number_code_remaining_verification_attempts, phone_number_code_is_verified, phone_number_code_error_message, name_value, name_is_valid, name_error_message, username_value, username_is_valid, username_error_message, password_value, password_is_valid, password_error_message
`

type CreateRegistrationFlowParams struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpireAt  time.Time
	CsrfToken string
	State     RegistrationFlowsState
}

func (q *Queries) CreateRegistrationFlow(ctx context.Context, arg *CreateRegistrationFlowParams) (RegistrationFlow, error) {
	row := q.db.QueryRow(ctx, createRegistrationFlow,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ExpireAt,
		arg.CsrfToken,
		arg.State,
	)
	var i RegistrationFlow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
		&i.ExpireAt,
		&i.CsrfToken,
		&i.State,
		&i.EmailValue,
		&i.EmailIsValid,
		&i.EmailErrorMessage,
		&i.EmailCode,
		&i.EmailCodeExpiry,
		&i.EmailCodeIncorrectVerificationAttempts,
		&i.EmailCodeRemainingVerificationAttempts,
		&i.EmailCodeIsVerified,
		&i.EmailCodeErrorMessage,
		&i.PhoneNumberValue,
		&i.PhoneNumberIsValid,
		&i.PhoneNumberErrorMessage,
		&i.PhoneNumberCode,
		&i.PhoneNumberCodeExpiry,
		&i.PhoneNumberCodeIncorrectVerificationAttempts,
		&i.PhoneNumberCodeRemainingVerificationAttempts,
		&i.PhoneNumberCodeIsVerified,
		&i.PhoneNumberCodeErrorMessage,
		&i.NameValue,
		&i.NameIsValid,
		&i.NameErrorMessage,
		&i.UsernameValue,
		&i.UsernameIsValid,
		&i.UsernameErrorMessage,
		&i.PasswordValue,
		&i.PasswordIsValid,
		&i.PasswordErrorMessage,
	)
	return i, err
}

const getRegistrationFlow = `-- name: GetRegistrationFlow :one
SELECT
    id, created_at, updated_at, version, expire_at, csrf_token, state, email_value, email_is_valid, email_error_message, email_code, email_code_expiry, email_code_incorrect_verification_attempts, email_code_remaining_verification_attempts, email_code_is_verified, email_code_error_message, phone_number_value, phone_number_is_valid, phone_number_error_message, phone_number_code, phone_number_code_expiry, phone_number_code_incorrect_verification_attempts, phone_number_code_remaining_verification_attempts, phone_number_code_is_verified, phone_number_code_error_message, name_value, name_is_valid, name_error_message, username_value, username_is_valid, username_error_message, password_value, password_is_valid, password_error_message
FROM
    "registration_flows"
WHERE
    "id" = $1
`

func (q *Queries) GetRegistrationFlow(ctx context.Context, id int64) (RegistrationFlow, error) {
	row := q.db.QueryRow(ctx, getRegistrationFlow, id)
	var i RegistrationFlow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
		&i.ExpireAt,
		&i.CsrfToken,
		&i.State,
		&i.EmailValue,
		&i.EmailIsValid,
		&i.EmailErrorMessage,
		&i.EmailCode,
		&i.EmailCodeExpiry,
		&i.EmailCodeIncorrectVerificationAttempts,
		&i.EmailCodeRemainingVerificationAttempts,
		&i.EmailCodeIsVerified,
		&i.EmailCodeErrorMessage,
		&i.PhoneNumberValue,
		&i.PhoneNumberIsValid,
		&i.PhoneNumberErrorMessage,
		&i.PhoneNumberCode,
		&i.PhoneNumberCodeExpiry,
		&i.PhoneNumberCodeIncorrectVerificationAttempts,
		&i.PhoneNumberCodeRemainingVerificationAttempts,
		&i.PhoneNumberCodeIsVerified,
		&i.PhoneNumberCodeErrorMessage,
		&i.NameValue,
		&i.NameIsValid,
		&i.NameErrorMessage,
		&i.UsernameValue,
		&i.UsernameIsValid,
		&i.UsernameErrorMessage,
		&i.PasswordValue,
		&i.PasswordIsValid,
		&i.PasswordErrorMessage,
	)
	return i, err
}