package main

import (
	"gotodo/app/config"
	"gotodo/app/services"
	"gotodo/http/handlers"
	"gotodo/http/server"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	servicesContainer := services.BuildServices(cfg)
	toDoHandler := todohandler.New(servicesContainer.ToDoService, servicesContainer.JwtService)
	err = server.Start(toDoHandler, cfg.ServerPort)

	if err != nil {
		log.Fatal(err)
	}
}
