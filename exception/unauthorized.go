package exception

import "fmt"

type UnauthorizedError struct {
	msg string
}

func NewUnauthorizedErrorError(msg string) *UnauthorizedError {
	return &UnauthorizedError{
		msg: msg,
	}
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("unauthorized: %s", e.msg)
}
