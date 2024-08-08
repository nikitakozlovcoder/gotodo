package services

import (
	"gotodo/app/config"
	"gotodo/app/repositories"
	"gotodo/app/repositories/connection"
)

type Services struct {
	ToDoService *ToDoService
	JwtService  *JwtService
	UserService *UserService
}

func BuildServices(cfg *config.Config) *Services {
	dbConnector := connection.NewDbConnector(cfg.DatabaseConnectionString)
	todoRepository := repositories.NewToDoRepository(dbConnector)
	jwtService := NewJwtService(cfg.JwtKey)

	return &Services{
		ToDoService: NewToDoService(todoRepository),
		JwtService:  jwtService,
		UserService: NewUserService(jwtService),
	}
}
