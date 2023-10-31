package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-3/service"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) RegisterNewUser(ctx *gin.Context) {
	user := entity.User{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		apiErr := errs.NewBadRequest("invalid json body")
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	response, err := h.userService.RegisterNewUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.NewInternalServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
