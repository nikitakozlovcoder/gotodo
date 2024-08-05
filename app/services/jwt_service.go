package services

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strconv"
	"time"
)

type JwtService struct {
	key string
}

func NewJwtService(key string) *JwtService {
	return &JwtService{key: key}
}

func (service *JwtService) CreateJwt(userId int64) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 24)},
		Subject:   strconv.FormatInt(userId, 10),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(service.key))
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
}

func (service *JwtService) ParseJwt(tokenString string) (int64, error) {
	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(service.key), nil
	})

	if err != nil {
		log.Println(err)
		return 0, err
	}

	subject, err := jwtToken.Claims.GetSubject()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := strconv.ParseInt(subject, 10, 64)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil
}
