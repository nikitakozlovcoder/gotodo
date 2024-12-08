package handlers

import (
	"gotodo/app/models/requests"
	"gotodo/app/services"
	"gotodo/http/middlewares"
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
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
	router.DELETE("/:id", h.deleteToDo)
}

func (h *ToDoHandler) listToDos(ctx *gin.Context) {
	todos, err := h.todoService.GetAll(ctx)
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

	id, err := h.todoService.SaveToDo(ctx, request)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(200, gin.H{"id": id})
}

func (h *ToDoHandler) deleteToDo(ctx *gin.Context) {
	var id, err = strconv.ParseInt(ctx.Param("id"), 10, 64) 
	if err != nil {
		log.Print(err)
		ctx.Status(http.StatusBadRequest)
		return
	}
	
	err = h.todoService.DeleteToDo(ctx, id)
	if err != nil {
		log.Print(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(200)
}
