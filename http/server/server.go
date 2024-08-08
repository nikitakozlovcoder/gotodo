package server

import (
	"github.com/gin-gonic/gin"
	"gotodo/http/handlers"
)

func Start(todoHandler *handlers.ToDoHandler, userHandler *handlers.UserHandler, port string) error {
	r := gin.Default()
	todoHandler.Init(r.Group("/todo"))
	userHandler.Init(r.Group("/user"))

	return r.Run(port)
}
