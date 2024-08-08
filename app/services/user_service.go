package services

import apperrors "gotodo/app/errors"

type IUserService interface {
	Login(userName string, password string) (string, error)
}

type UserService struct {
	jwtService IJwtService
}

func NewUserService(jwtService IJwtService) *UserService {
	return &UserService{jwtService: jwtService}
}

func (userService *UserService) Login(userName string, password string) (string, error) {
	return "", apperrors.InvalidUserCredentials
}
