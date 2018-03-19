package app

import "github.com/prokosna/medusa_synapse/domain"

type Medusa struct {
	publisher domain.Publisher
	validator domain.Validator
}

func NewMedusa(
	publisher domain.Publisher,
	validator domain.Validator) *Medusa {
	return &Medusa{
		publisher: publisher,
		validator: validator,
	}
}

func (m *Medusa) SendFrame(key string, img domain.Image) error {
	err := m.validator.Validate(img)
	if err != nil {
		return err
	}
	err = m.publisher.Publish(key, img)
	if err != nil {
		return err
	}
	return nil
}
