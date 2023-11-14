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
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response, err := ch.categoryService.CreateCategory(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.NewInternalServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ch *categoryHandler) GetCategoryWithTask(ctx *gin.Context) {
	response, err := ch.categoryService.GetCategoryWithTask()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ch *categoryHandler) UpdateCategory(ctx *gin.Context) {
	var request dto.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
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
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}
