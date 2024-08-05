package services

import (
	"gotodo/app/config"
	"gotodo/app/repositories"
	"gotodo/app/repositories/connection"
)

type Services struct {
	ToDoService *ToDoService
}

func BuildServices(cfg *config.Config) *Services {
	dbConnector := connection.NewDbConnector(cfg.DatabaseConnectionString)
	todoRepository := repositories.NewToDoRepository(dbConnector)

	return &Services{
		ToDoService: NewToDoService(todoRepository),
	}
}
