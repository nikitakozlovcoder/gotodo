package repositories

import "gotodo/app/repositories/connection"

type IUserRepository interface {
	GetUserPasswordHashByUserName(userName string, password string) (string, error)
}

type UserRepository struct {
	connection connection.Executor
}

func NewUserRepository(connection connection.Executor) *UserRepository {
	return &UserRepository{connection: connection}
}

func (repo *UserRepository) GetUserPasswordHashByUserName(userName string, password string) (string, error) {
	return "", nil
}
