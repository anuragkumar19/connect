package serviceerrors

type ServiceError interface {
	error
	Type() ErrType
	InternalError() error
}

func NewInternalError(err error) error {
	return New(TypeInternal, "Internal server error", err)
}
