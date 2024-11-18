package users

import "time"

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
	Token               string
	ResendCodeAfterTime time.Time
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
