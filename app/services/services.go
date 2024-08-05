package services

import (
	"gotodo/app/config"
	"gotodo/app/repositories"
	"gotodo/app/repositories/connection"
)

type Services struct {
	ToDoService *ToDoService
	JwtService  *JwtService
}

func BuildServices(cfg *config.Config) *Services {
	dbConnector := connection.NewDbConnector(cfg.DatabaseConnectionString)
	todoRepository := repositories.NewToDoRepository(dbConnector)

	return &Services{
		ToDoService: NewToDoService(todoRepository),
		JwtService:  NewJwtService(cfg.JwtKey),
	}
}
