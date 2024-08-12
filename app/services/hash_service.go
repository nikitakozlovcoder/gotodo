package services

type IHashService interface {
	Hash(key string) string
}

type HashSha256Service struct{}

func NewHashService() *HashSha256Service {
	return &HashSha256Service{}
}

func (HashSha256Service) Hash(key string) string {
	return ""
}
