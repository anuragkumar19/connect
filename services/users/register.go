package users

import (
	"context"
	"errors"

	"github.com/anuragkumar19/connect/database"
	"github.com/anuragkumar19/connect/pkg/argon2id"
	"github.com/anuragkumar19/connect/services/serviceerrors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/oklog/ulid/v2"
)

func (s *Users) Register(ctx context.Context, cmd RegisterCmd) (RegisterResult, error) {
	var result RegisterResult
	if err := cmd.Transform().Validate(ctx); err != nil {
		return result, err
	}

	if err := pgx.BeginFunc(ctx, s.db, func(tx pgx.Tx) error {
		queries := s.queries.WithTx(tx)
		available, err := queries.IsUsernameAvailable(ctx, cmd.Username)
		if err != nil {
			return serviceerrors.NewInternalError(err)
		}
		if !available.Valid {
			return serviceerrors.NewInternalError(errors.New("expected IsUsernameAvailable query to return non null value"))
		}
		if !available.Bool {
			return serviceerrors.NewFieldError("Username", "Username already taken")
		}

		userId := ulid.Make()
		emailId := ulid.Make()
		phoneNumberId := ulid.Make()

		if err := queries.CreateUser(ctx, &database.CreateUserParams{
			ID:           userId,
			IsRegistered: false,
			Name:         cmd.Name,
			Username:     cmd.Username,
			PrimaryEmailID: database.ULIDValue{
				Valid: cmd.Method == Email,
				ULID:  emailId,
			},
			PrimaryPhoneNumberID: database.ULIDValue{
				Valid: cmd.Method == PhoneNumber,
				ULID:  phoneNumberId,
			},
		}); err != nil {
			return serviceerrors.NewInternalError(err)
		}

		h, err := argon2id.CreateHash(cmd.Password, argon2id.DefaultParams) // TODO: use passwords service, we can also use new db domain hashed
		if err != nil {
			return serviceerrors.NewInternalError(err)
		}
		if err := queries.CreatePassword(ctx, &database.CreatePasswordParams{
			ID:       ulid.Make(),
			IsActive: true,
			UserID:   userId,
			Value:    h,
		}); err != nil {
			return serviceerrors.NewInternalError(err)
		}

		if cmd.Method == Email {
			available, err := queries.IsEmailAvailable(ctx, cmd.Email)
			if err != nil {
				return serviceerrors.NewInternalError(err)
			}
			if !available {
				return serviceerrors.NewFieldError("Email", "A account is already linked to provided email")
			}

			if err := queries.CreateEmail(ctx, &database.CreateEmailParams{
				ID:         emailId,
				UserID:     userId,
				Value:      cmd.Email,
				IsPrimary:  true,
				IsVerified: false,
				LastVerifiedAt: pgtype.Timestamptz{
					Valid: false,
				},
			}); err != nil {
				return serviceerrors.NewInternalError(err)
			}

			// otp+send_email
		}

		if cmd.Method == PhoneNumber {
			available, err := queries.IsPhoneNumberAvailable(ctx, cmd.PhoneNumber)
			if err != nil {
				return serviceerrors.NewInternalError(err)
			}
			if !available {
				return serviceerrors.NewFieldError("PhoneNumber", "A account is already linked to provided phoneNumber")
			}

			if err := queries.CreatePhoneNumber(ctx, &database.CreatePhoneNumberParams{
				ID:         phoneNumberId,
				UserID:     userId,
				Value:      cmd.PhoneNumber,
				IsPrimary:  true,
				IsVerified: false,
				LastVerifiedAt: pgtype.Timestamptz{
					Valid: false,
				},
			}); err != nil {
				return serviceerrors.NewInternalError(err)
			}
		}

		return nil
	}); err != nil {
		if _, ok := err.(serviceerrors.ServiceError); ok {
			return result, err
		}

		return result, serviceerrors.NewInternalError(err)
	}

	return result, nil
}

func (s *Users) ResendVerificationCode(ctx context.Context, cmd ResendVerificationCodeCmd) (ResendVerificationCodeResult, error) {
	return ResendVerificationCodeResult{}, nil
}

func (s *Users) Verify(ctx context.Context, cmd VerifyCmd) error {
	return nil
}
