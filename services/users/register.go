package users

import "context"

func (s *Users) Register(ctx context.Context, cmd RegisterCmd) (RegisterResult, error) {
	return RegisterResult{}, nil
}

func (s *Users) ResendVerificationCode(ctx context.Context, cmd ResendVerificationCodeCmd) (ResendVerificationCodeResult, error) {
	return ResendVerificationCodeResult{}, nil
}

func (s *Users) Verify(ctx context.Context, cmd VerifyCmd) error {
	return nil
}
