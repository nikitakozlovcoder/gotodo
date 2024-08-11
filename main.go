package main

import (
	"gotodo/app/config"
	"gotodo/app/repositories/connection"
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

	dbConnect, err := connection.NewDbConnector(cfg.DatabaseConnectionString).DbConnect()
	defer dbConnect.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	servicesContainer := services.BuildServices(cfg, dbConnect)
	jwt, _ := servicesContainer.JwtService.CreateJwt(1)
	log.Println(jwt)
	authMiddleware := middlewares.NewAuthMiddleware(servicesContainer.JwtService)
	toDoHandler := handlers.NewToDoHandler(servicesContainer.ToDoService, authMiddleware)
	userHandler := handlers.NewUserHandler(servicesContainer.UserService)
	err = server.Start(toDoHandler, userHandler, cfg.ServerPort)
	if err != nil {
		log.Fatal(err)
	}
}
