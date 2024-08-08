package handlers

import (
	"github.com/gin-gonic/gin"
	"gotodo/app/models/requests"
	"gotodo/app/services"
	"gotodo/http/middlewares"
	"log"
	"net/http"
)

type ToDoHandler struct {
	todoService    services.IToDoService
	authMiddleware *middlewares.AuthMiddleware
}

func NewToDoHandler(toDoService services.IToDoService,
	authMiddleware *middlewares.AuthMiddleware) *ToDoHandler {
	return &ToDoHandler{
		todoService:    toDoService,
		authMiddleware: authMiddleware,
	}
}

func (h *ToDoHandler) Init(router *gin.RouterGroup) {
	router.GET("/list", h.listToDos)
	router.POST("", h.authMiddleware.Handle(), h.addToDo)
}

func (h *ToDoHandler) listToDos(ctx *gin.Context) {
	todos, err := h.todoService.GetAll()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(200, todos)
}

func (h *ToDoHandler) addToDo(ctx *gin.Context) {
	var request requests.NewToDoRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		log.Print(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	id, err := h.todoService.SaveToDo(request)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(200, gin.H{"id": id})
}
