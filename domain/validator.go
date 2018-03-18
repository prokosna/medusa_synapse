package domain

type Validator interface {
	validate(message string) error
}
