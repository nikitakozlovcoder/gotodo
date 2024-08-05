package todohandler

import (
	"github.com/gin-gonic/gin"
	"gotodo/app/models/requests"
	"gotodo/app/services"
	"gotodo/http/middlewares"
	"log"
	"net/http"
)

type IToDoHandler interface {
	Init(ctx *gin.RouterGroup)
}

type ToDoHandler struct {
	todoService services.IToDoService
	jwtService  *services.JwtService
}

func New(toDoService services.IToDoService, jwtService *services.JwtService) *ToDoHandler {
	return &ToDoHandler{
		todoService: toDoService,
		jwtService:  jwtService,
	}
}

func (h *ToDoHandler) Init(router *gin.RouterGroup) {
	router.GET("/list", h.listToDos)
	router.POST("", middlewares.AuthMiddleware(h.jwtService), h.addToDo)
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
