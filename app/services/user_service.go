package services

import (
	"context"
	"gotodo/app/apperrors"
	"gotodo/app/repositories"
)

type IUserService interface {
	Login(ctx context.Context, login string, password string) (string, error)
}

type UserService struct {
	jwtService     IJwtService
	userRepository repositories.IUserRepository
	hashService    IHashService
}

func NewUserService(jwtService IJwtService,
	userRepository repositories.IUserRepository,
	hashService IHashService) *UserService {
	return &UserService{jwtService: jwtService, userRepository: userRepository, hashService: hashService}
}

func (userService *UserService) Login(ctx context.Context, login string, password string) (string, error) {
	user, err := userService.userRepository.GetUserByLogin(ctx, login)
	if err != nil {
		return "", err
	}

	providedPasswordHash := userService.hashService.Hash(password)
	if providedPasswordHash != user.PasswordHash {
		return "", apperrors.InvalidUserCredentials
	}

	jwt, err := userService.jwtService.CreateJwt(user.Id)
	if err != nil {
		return "", err
	}

	return jwt, apperrors.InvalidUserCredentials
}
