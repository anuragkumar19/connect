package serviceerrors

type messageErr struct {
	msg string
	err error
	tp  ErrType
}

var _ ServiceError = (*messageErr)(nil)

func (e *messageErr) Error() string {
	return e.msg
}

func (e *messageErr) Type() ErrType {
	return e.tp
}

func (e *messageErr) InternalError() error {
	return e.err
}

func New(tp ErrType, message string, internalErr error) error {
	return &messageErr{
		msg: message,
		err: internalErr,
		tp:  tp,
	}
}
