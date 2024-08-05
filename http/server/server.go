package server

import (
	"github.com/gin-gonic/gin"
	"gotodo/http/handlers"
)

func Start(todoHandler todohandler.IToDoHandler, port string) error {
	r := gin.Default()
	todoHandler.Init(r.Group("/todo"))

	return r.Run(port)
}
