package users

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	_ "github.com/nyaruka/phonenumbers"
)

var NameRules = []validation.Rule{
	validation.Length(2, 30),
}

var UsernameRules = []validation.Rule{
	validation.Length(4, 30),
	validation.Match(regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]*$")).Error("can start with alphabets only and contains only alphabets, numbers and underscores"),
}

var PasswordRules = []validation.Rule{
	validation.Length(8, 250),
}

var EmailRules = []validation.Rule{
	is.Email,
}

var PhoneNumberRules = []validation.Rule{}
