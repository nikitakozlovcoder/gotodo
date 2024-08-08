package main

import (
	"gotodo/app/config"
	"gotodo/app/services"
	"gotodo/http/handlers"
	"gotodo/http/middlewares"
	"gotodo/http/server"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	servicesContainer := services.BuildServices(cfg)
	authMiddleware := middlewares.NewAuthMiddleware(servicesContainer.JwtService)
	toDoHandler := handlers.NewToDoHandler(servicesContainer.ToDoService, authMiddleware)
	userHandler := handlers.NewUserHandler(servicesContainer.UserService)
	err = server.Start(toDoHandler, userHandler, cfg.ServerPort)

	if err != nil {
		log.Fatal(err)
	}
}
