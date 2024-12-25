package server

import "fmt"

var (
	ERRINVALIDPARAMS = NewError(10001, "invalid params")
)

type Error struct {
	Code    int
	Message string
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func (e *Error) Error() error {
	return fmt.Errorf(e.Message)
}
