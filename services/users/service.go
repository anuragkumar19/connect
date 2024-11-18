package users

import (
	"context"
)

type Service interface {
	Register(context.Context, RegisterCmd) (RegisterResult, error)
	ResendVerificationCode(context.Context, ResendVerificationCodeCmd) (ResendVerificationCodeResult, error)
	Verify(context.Context, VerifyCmd) error
}
