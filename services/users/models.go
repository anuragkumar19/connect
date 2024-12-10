package users

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/anuragkumar19/connect/services/serviceerrors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type RegisterMethod int

const (
	Email RegisterMethod = iota + 1
	PhoneNumber
)

type RegisterCmd struct {
	Name        string
	Username    string
	Method      RegisterMethod
	Email       string
	PhoneNumber string
	Password    string
}

type RegisterResult struct {
	Token string
}

type ResendVerificationCodeCmd struct {
	Token string
}

type ResendVerificationCodeResult struct {
	ResendCodeAfterTime time.Time
}

type VerifyCmd struct {
	Token string
	Code  string
}

func (s *RegisterCmd) Transform() *RegisterCmd {
	s.Email = strings.TrimSpace(s.Email)
	s.Name = strings.TrimSpace(s.Name)
	s.PhoneNumber = strings.TrimSpace(s.PhoneNumber)
	return s
}

// TODO: use ruled defined in validation.go
func (s *RegisterCmd) Validate(ctx context.Context) error {
	if err := validation.ValidateStructWithContext(
		ctx,
		&s,
		validation.Field(&s.Email, validation.Required.When(s.Method == Email), is.EmailFormat),
		validation.Field(&s.Method, validation.Required, validation.In(Email, PhoneNumber)),
		validation.Field(&s.Name, validation.Required, validation.Length(2, 30)),
		validation.Field(&s.Password, validation.Required, validation.Length(8, 250)),
		validation.Field(&s.PhoneNumber, validation.Required.When(s.Method == PhoneNumber)),
		validation.Field(&s.Username, validation.Required, validation.Length(4, 30), validation.Match(regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]*$")).Error("can start with alphabets only and contains only alphabets, numbers and underscores")),
	); err != nil {
		if _, ok := err.(validation.InternalError); ok {
			return serviceerrors.NewInternalError(err)
		}
		return serviceerrors.NewInvalidArgumentFromOzzoErrors(err)
	}
	return nil
}
