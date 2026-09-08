package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hacktiv8-ks07-g04/final-project-3/dto"
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

// ISSUE: Receiver naming is inconsistent — this file uses both `h` (RegisterNewUser)
// and `uh` (LoginUser, UpdateUser, DeleteUser) for the same *userHandler type.
// Go convention: pick one receiver name and use it consistently throughout.

func (h *userHandler) RegisterNewUser(ctx *gin.Context) {
	var newRequest dto.RegisterRequest

	err := ctx.ShouldBindJSON(&newRequest)
	if err != nil {
		apiErr := errs.NewBadRequest("invalid json body")
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	response, err := h.userService.CreateNewUser(&newRequest)
	if err != nil {
		// ISSUE: Wraps the original error into a generic InternalServerError,
		// discarding the specific error type/status from the service layer.
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

// update user without using userid params
func (uh *userHandler) UpdateUser(ctx *gin.Context) {
	var newRequest dto.UpdateUserRequest

	if err := ctx.ShouldBindJSON(&newRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("Invalid Request Body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := uh.userService.UpdateUser(newRequest)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

// delete user
func (uh *userHandler) DeleteUser(ctx *gin.Context) {
	// ISSUE: ctx.MustGet + unchecked type assertion .(entity.User). If "userData"
	// is missing from context or stored as a different type, this panics and crashes
	// the server. Use ctx.Get() with an ok check.
	user := ctx.MustGet("userData").(entity.User)

	response, err := uh.userService.DeleteUser(user.ID)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	// ISSUE: Uses response.Status (DeleteUserResponse uses `Status` field) while
	// other handlers use response.StatusCode — inconsistent field naming across DTOs.
	ctx.JSON(response.Status, response)
}
