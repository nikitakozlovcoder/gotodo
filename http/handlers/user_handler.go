package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	apperrors "gotodo/app/errors"
	"gotodo/app/models/requests"
	"gotodo/app/services"
	"log"
	"net/http"
)

type UserHandler struct {
	userService services.IUserService
}

func NewUserHandler(userService services.IUserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Authenticate() {}

func (h *UserHandler) Init(router *gin.RouterGroup) {
	router.POST("/login", h.login)
}

func (h *UserHandler) login(ctx *gin.Context) {
	var request requests.UserLoginRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		log.Println(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	jwt, err := h.userService.Login(request.UserName, request.Password)
	if err != nil {
		log.Println(err)
		if errors.Is(err, apperrors.InvalidUserCredentials) {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
