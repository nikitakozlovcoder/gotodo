package services

import (
	"crypto/sha256"
	"encoding/base64"
)

type IHashService interface {
	Hash(key string) string
}

type HashSha256Service struct{}

func NewHashService() *HashSha256Service {
	return &HashSha256Service{}
}

func (HashSha256Service) Hash(key string) string {
	sha := sha256.Sum256([]byte(key))
	return base64.URLEncoding.EncodeToString(sha[:])
}
