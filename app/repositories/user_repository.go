package repositories

import (
	"context"
	"database/sql"
	"errors"
	"gotodo/app/repositories/connection"
)

type IUserRepository interface {
	GetUserByLogin(ctx context.Context, login string) (*struct {
		Id           int64
		PasswordHash string
	}, error)
}

type UserRepository struct {
	connection connection.Executor
}

func NewUserRepository(connection connection.Executor) *UserRepository {
	return &UserRepository{connection: connection}
}

func (repo *UserRepository) GetUserByLogin(ctx context.Context, login string) (*struct {
	Id           int64
	PasswordHash string
}, error) {
	row := repo.connection.QueryRowContext(ctx, "SELECT id, password_hash FROM Users WHERE login = $1", login)
	if err := row.Err(); err != nil {
		return nil, err
	}

	user := struct {
		Id           int64
		PasswordHash string
	}{}

	err := row.Scan(&user.Id, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
