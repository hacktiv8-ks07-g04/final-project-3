package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hacktiv8-ks07-g04/final-project-3/dto"
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
	var newRequest dto.RegisterRequest

	err := ctx.ShouldBindJSON(&newRequest)
	if err != nil {
		apiErr := errs.NewBadRequest("invalid json body")
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	response, err := h.userService.CreateNewUser(newRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.NewInternalServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (uh *userHandler) LoginUser(ctx *gin.Context) {
	var newUserRequest dto.LoginRequest

	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("Invalid Request Body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := uh.userService.LoginUser(newUserRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}
