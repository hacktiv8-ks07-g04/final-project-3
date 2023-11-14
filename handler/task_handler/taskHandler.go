package task_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hacktiv8-ks07-g04/final-project-3/dto"
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-3/service"
)

type taskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) *taskHandler {
	return &taskHandler{taskService: taskService}
}

func (h *taskHandler) CreateNewTask(ctx *gin.Context) {
	var newRequest dto.NewTaskRequest
	user := ctx.MustGet("userData").(entity.User)

	err := ctx.ShouldBindJSON(&newRequest)
	if err != nil {
		apiErr := errs.NewUnprocessibleEntityError("invalid json body")
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	response, err := h.taskService.CreateNewTask(user.UserID, &newRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.NewInternalServerError("Error when trying to create new task"))
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
