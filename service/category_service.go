package service

import (
	"net/http"

	"github.com/hacktiv8-ks07-g04/final-project-3/dto"
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/helpers"
	"github.com/hacktiv8-ks07-g04/final-project-3/repository/category_repository"
)

type CategoryService interface {
	CreateCategory(payload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.MessageErr)
	GetCategoryWithTask() (*dto.CategoryListResponse, errs.MessageErr)
	UpdateCategory(categoryID uint, payload *dto.UpdateCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr)
	DeleteCategory(categoryID uint) (*dto.DeleteCategoryResponse, errs.MessageErr)
}

type categoryService struct {
	categoryRepo category_repository.Repository
}

func NewCategoryService(categoryRepo category_repository.Repository) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (c *categoryService) CreateCategory(payload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)
	if err != nil {
		return nil, err
	}

	category := entity.Category{
		Type: payload.Type,
	}

	err = c.categoryRepo.CreateCategory(&category)
	if err != nil {
		return nil, err
	}

	response := dto.NewCategoryResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully created new category",
		Data: dto.NewCategoryDataResponse{
			CategoryID: category.ID,
			Type:       category.Type,
			CreatedAt:  category.CreatedAt,
		},
	}

	return &response, nil
}

func (c *categoryService) GetCategoryWithTask() (*dto.CategoryListResponse, errs.MessageErr) {
	allCategories, err := c.categoryRepo.GetCategoryWithTask()
	if err != nil {
		return nil, errs.NewInternalServerError("Error when trying to get data")
	}

	categories := []dto.Category{}
	for _, v := range allCategories {
		tasks := []dto.Task{}
		for _, t := range v.Task {
			task := dto.Task{
				ID:          t.ID,
				Title:       t.Title,
				Description: t.Description,
				UserID:      t.UserID,
				CategoryID:  t.CategoryID,
				CreatedAt:   t.CreatedAt,
				UpdatedAt:   t.UpdatedAt,
			}
			tasks = append(tasks, task)
		}

		category := dto.Category{
			CategoryID: v.ID,
			Type:       v.Type,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
			Tasks:      tasks,
		}
		categories = append(categories, category)
	}

	response := dto.CategoryListResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully get all categories",
		Data:       categories,
	}

	return &response, nil
}

// update category
func (c *categoryService) UpdateCategory(categoryID uint, payload *dto.UpdateCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)
	if err != nil {
		return nil, err
	}

	category := &entity.Category{
		Type: payload.Type,
	}

	result, err := c.categoryRepo.UpdateCategory(categoryID, category)
	if err != nil {
		return nil, err
	}

	response := dto.UpdateCategoryResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully updated category",
		Data: dto.UpdateCategoryDataResponse{
			CategoryID: result.ID,
			Type:       result.Type,
			UpdatedAt:  result.UpdatedAt,
		},
	}

	return &response, nil
}

// delete category
func (c *categoryService) DeleteCategory(categoryID uint) (*dto.DeleteCategoryResponse, errs.MessageErr) {
	err := c.categoryRepo.DeleteCategory(categoryID)
	if err != nil {
		return nil, err
	}

	response := dto.DeleteCategoryResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully deleted category",
	}

	return &response, nil
}
