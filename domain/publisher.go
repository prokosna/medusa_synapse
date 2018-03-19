package domain

type Publisher interface {
	Publish(key string, img Image) error
}
