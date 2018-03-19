package domain

type Validator interface {
	Validate(img Image) error
}
