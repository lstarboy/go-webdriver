package errors

type ExpectError struct {
	message string
}

func (e *ExpectError) Error() string {
	return e.message
}

func NewExpectError(msg string) error {
	return &ExpectError{message: msg}
}
