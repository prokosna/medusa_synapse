package domain

type Publisher interface {
	publish(message string) error
}
