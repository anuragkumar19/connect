package serviceerrors

import (
	"fmt"
	"slices"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type FieldError struct {
	field   string
	message string
}

func (e *FieldError) Field() string {
	return e.field
}

func (e *FieldError) Message() string {
	return e.message
}

type InvalidArgumentError struct {
	fieldErrs []FieldError
	err       error
	message   string
}

var _ ServiceError = (*InvalidArgumentError)(nil)

func NewInvalidArgumentError(errs []FieldError) error {
	return &InvalidArgumentError{
		fieldErrs: errs,
	}
}

func NewInvalidArgumentFromOzzoErrors(err error) error {
	if e, ok := err.(validation.InternalError); ok {
		return NewInternalError(e)
	}
	errs, ok := err.(validation.Errors)
	if !ok {
		return NewInternalError(fmt.Errorf("unexpected error received, expected validation.Errors: %w", err))
	}
	vErrs := make([]FieldError, 0, len(errs))

	for k, v := range errs {
		vErrs = append(vErrs, FieldError{
			field:   k,
			message: v.Error(),
		})
	}

	return &InvalidArgumentError{
		fieldErrs: vErrs,
	}
}

func NewFieldError(field string, message string) error {
	return &InvalidArgumentError{
		fieldErrs: []FieldError{
			{
				field:   field,
				message: message,
			},
		},
	}
}

func (e *InvalidArgumentError) Error() string {
	if e.message == "" {
		var s strings.Builder
		for i, vErr := range e.fieldErrs {
			if i != 0 {
				s.WriteString(", ")
			}
			s.WriteString(fmt.Sprintf("%s: %s", vErr.field, vErr.message))
			if i == len(e.fieldErrs)-1 {
				s.WriteString(".")
			}
		}
		e.message = s.String()
	}

	return e.message
}

func (e *InvalidArgumentError) Type() ErrType {
	return TypeInvalidArgument
}

func (e *InvalidArgumentError) InternalError() error {
	return e.err
}

func (e *InvalidArgumentError) Errors() []FieldError {
	return slices.Clone(e.fieldErrs)
}
