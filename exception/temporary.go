package exception

import "fmt"

type ServiceUnavailableError struct {
	msg string
}

func NewServiceUnavailableError(msg string) *ServiceUnavailableError {
	return &ServiceUnavailableError{
		msg: msg,
	}
}

func (e *ServiceUnavailableError) Error() string {
	return fmt.Sprintf("service unavailable: %s", e.msg)
}
