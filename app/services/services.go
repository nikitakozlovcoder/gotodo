package services

import (
	"gotodo/app/config"
	"gotodo/app/repositories"
	"gotodo/app/repositories/connection"
	"gotodo/app/repositories/transaction"
)

type Services struct {
	ToDoService *ToDoService
	JwtService  *JwtService
	UserService *UserService
}

func BuildServices(cfg *config.Config, connection *connection.DbConnection) *Services {
	transactionManager := transaction.NewManager(connection)
	todoRepository := repositories.NewToDoRepository(connection)
	jwtService := NewJwtService(cfg.JwtKey)

	return &Services{
		ToDoService: NewToDoService(todoRepository, transactionManager),
		JwtService:  jwtService,
		UserService: NewUserService(jwtService),
	}
}
