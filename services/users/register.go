package users

import (
	"context"
	"errors"
)

func (s *Users) Register(ctx context.Context, cmd RegisterCmd) (RegisterResult, error) {
	if err := cmd.Transform().Validate(); err != nil {
		return RegisterResult{}, err
	}

	exist, err := s.queries.IsUsernameAvailable(ctx, cmd.Username)
	if err != nil {
		return RegisterResult{}, err
	}
	if !exist.Bool {
		return RegisterResult{}, errors.New("username taken")
	}

	return RegisterResult{}, nil
}

func (s *Users) ResendVerificationCode(ctx context.Context, cmd ResendVerificationCodeCmd) (ResendVerificationCodeResult, error) {
	return ResendVerificationCodeResult{}, nil
}

func (s *Users) Verify(ctx context.Context, cmd VerifyCmd) error {
	return nil
}
