package exception

import "fmt"

type BadRequestError struct {
	msg string
}

func NewBadRequestError(msg string) *BadRequestError {
	return &BadRequestError{
		msg: msg,
	}
}

func (e *BadRequestError) Error() string {
	return fmt.Sprintf("bad request: %s", e.msg)
}
