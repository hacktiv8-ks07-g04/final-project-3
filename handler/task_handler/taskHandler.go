package task_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hacktiv8-ks07-g04/final-project-3/dto"
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/helpers"
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

	response, err := h.taskService.CreateNewTask(user.ID, &newRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.NewInternalServerError("Error when trying to create new task"))
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *taskHandler) GetTaskWithUser(ctx *gin.Context) {
	response, err := h.taskService.GetTaskWithUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.NewInternalServerError("Error when trying to get data"))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *taskHandler) UpdateTaskById(ctx *gin.Context) {
	var updateRequest dto.UpdateDetailTaskRequest
	user := ctx.MustGet("userData").(entity.User)

	taskId, err := helpers.GetParamId(ctx, "taskId")
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		apiErr := errs.NewUnprocessibleEntityError("invalid json body")
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	response, err := h.taskService.UpdateTaskById(taskId, user.ID, &updateRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.NewInternalServerError("Error when trying to update task title and description field"))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *taskHandler) UpdateTaskStatus(ctx *gin.Context) {
	var updateStatus dto.UpdateTaskStatusRequest
	user := ctx.MustGet("userData").(entity.User)

	id, err := helpers.GetParamId(ctx, "taskId")
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because id not found")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	if err := ctx.ShouldBindJSON(&updateStatus); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json body")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := h.taskService.UpdateTaskStatus(id, user.ID, &updateStatus)
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred when updating task status")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

func (h *taskHandler) UpdateTaskCategory(ctx *gin.Context) {
	var updateCategory dto.UpdateTaskCategoryRequest
	user := ctx.MustGet("userData").(entity.User)

	id, err := helpers.GetParamId(ctx, "taskId")
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because id not found")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	if err := ctx.ShouldBindJSON(&updateCategory); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json body")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := h.taskService.UpdateTaskCategory(id, user.ID, &updateCategory)
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred when updating task category")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

func (h *taskHandler) DeleteTaskById(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)

	id, err := helpers.GetParamId(ctx, "taskId")
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because id not found")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := h.taskService.DeleteTaskById(id, user.ID)
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred when deleting task")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	ctx.JSON(result.StatusCode, result)
}
