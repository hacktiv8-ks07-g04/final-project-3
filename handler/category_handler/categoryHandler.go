package category_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hacktiv8-ks07-g04/final-project-3/dto"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/helpers"
	"github.com/hacktiv8-ks07-g04/final-project-3/service"
)

type categoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *categoryHandler {
	return &categoryHandler{categoryService: categoryService}
}

func (ch *categoryHandler) CreateCategory(ctx *gin.Context) {
	var request dto.NewCategoryRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		// ISSUE: Returns a raw error string (`err.Error()`) instead of the structured
		// errs.ErrorData JSON object used by other handlers. Inconsistent error format.
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response, err := ch.categoryService.CreateCategory(&request)
	if err != nil {
		// ISSUE: Wraps the original service error into a generic InternalServerError,
		// losing the specific error type/status.
		ctx.JSON(http.StatusInternalServerError, errs.NewInternalServerError(err.Error()))
		return
	}

	// ISSUE: Returns 200 OK for a creation. REST convention is 201 Created
	// (compare: user & task create handlers correctly return 201).
	ctx.JSON(http.StatusOK, response)
}

func (ch *categoryHandler) GetCategoryWithTask(ctx *gin.Context) {
	response, err := ch.categoryService.GetCategoryWithTask()
	if err != nil {
		// ISSUE: Returns raw error string instead of structured errs JSON object.
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ch *categoryHandler) UpdateCategory(ctx *gin.Context) {
	var request dto.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		// ISSUE: Raw error string instead of structured error.
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	categoryID, err := helpers.GetParamId(ctx, "categoryId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errs.NewBadRequest("invalid parameter id"))
		return
	}

	response, err := ch.categoryService.UpdateCategory(categoryID, &request)
	if err != nil {
		// ISSUE: Raw error string instead of structured error. Also always returns 500,
		// so a "not found" error from the repo surfaces as 500 instead of 404.
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ch *categoryHandler) DeleteCategory(ctx *gin.Context) {
	categoryID, err := helpers.GetParamId(ctx, "categoryId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errs.NewBadRequest("invalid parameter id"))
		return
	}

	response, err := ch.categoryService.DeleteCategory(categoryID)
	if err != nil {
		// ISSUE: Raw error string; also always 500, so "not found" becomes 500 not 404.
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}
