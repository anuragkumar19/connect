package users

import (
	"context"
	"errors"
	"time"

	"github.com/anuragkumar19/connect/database"
	"github.com/anuragkumar19/connect/pkg/crypto"
	"github.com/anuragkumar19/connect/services/serviceerrors"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

var (
	otpLength   = 6
	otpLifetime = time.Hour
)

func (s *Users) Register(ctx context.Context, cmd RegisterCmd) (RegisterResult, error) {
	var result RegisterResult
	if err := cmd.Transform().Validate(ctx); err != nil {
		return result, err
	}

	if err := s.store.BeginFunc(ctx, func(store database.Store) error {
		available, err := store.IsUsernameAvailable(ctx, cmd.Username)
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

		token, err := crypto.RandomToken()
		if err != nil {
			return serviceerrors.NewInternalError(err)
		}
		result.Token = token

		if err := store.CreateUser(ctx, &database.CreateUserParams{
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

		hashedPass, err := crypto.GenerateFromPassword(cmd.Password)
		if err != nil {
			return serviceerrors.NewInternalError(err)
		}
		if err := store.CreatePassword(ctx, &database.CreatePasswordParams{
			ID:       ulid.Make(),
			IsActive: true,
			UserID:   userId,
			Value:    hashedPass,
		}); err != nil {
			return serviceerrors.NewInternalError(err)
		}

		if cmd.Method == Email {
			available, err := store.IsEmailAvailable(ctx, cmd.Email)
			if err != nil {
				return serviceerrors.NewInternalError(err)
			}
			if !available {
				return serviceerrors.NewFieldError("Email", "A account is already linked to provided email")
			}

			if err := store.CreateEmail(ctx, &database.CreateEmailParams{
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

			opt, err := crypto.RandomNumber(otpLength)
			if err != nil {
				return serviceerrors.NewInternalError(err)
			}
			hashedOtp, err := crypto.GenerateFromPassword(opt)
			if err != nil {
				return serviceerrors.NewInternalError(err)
			}
			expiry := time.Now().Add(otpLifetime)
			if err := store.CreateVerificationToken(ctx, &database.CreateVerificationTokenParams{
				Token:    token,
				ID:       ulid.Make(),
				ExpireAt: expiry,
				UserID:   userId,
				EmailID: database.ULIDValue{
					Valid: true,
					ULID:  emailId,
				},
				PhoneNumberID: database.ULIDValue{
					Valid: false,
				},
				Otp: hashedOtp,
			}); err != nil {
				return serviceerrors.NewInternalError(err)
			}

			log.Info().Str("otp", opt).Str("token", token).Str("email", cmd.Email).Msg("email send") // TODO: remove
			// TODO: s.mailer.SendEmail(ctx, &mailerv1.SendEmail{})

			bucket := s.rateLimiter.UserTriggeredEmailBucket(ctx, cmd.Email)
			if err := bucket.Consume(ctx); err != nil {
				return serviceerrors.NewInternalError(err)
			}
		}

		if cmd.Method == PhoneNumber {
			return serviceerrors.New(serviceerrors.TypeUnimplemented, "Phone number registration not allowed yet", nil)
			// available, err := store.IsPhoneNumberAvailable(ctx, cmd.PhoneNumber)
			// if err != nil {
			// 	return serviceerrors.NewInternalError(err)
			// }
			// if !available {
			// 	return serviceerrors.NewFieldError("PhoneNumber", "A account is already linked to provided phoneNumber")
			// }

			// if err := store.CreatePhoneNumber(ctx, &database.CreatePhoneNumberParams{
			// 	ID:         phoneNumberId,
			// 	UserID:     userId,
			// 	Value:      cmd.PhoneNumber,
			// 	IsPrimary:  true,
			// 	IsVerified: false,
			// 	LastVerifiedAt: pgtype.Timestamptz{
			// 		Valid: false,
			// 	},
			// }); err != nil {
			// 	return serviceerrors.NewInternalError(err)
			// }
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
