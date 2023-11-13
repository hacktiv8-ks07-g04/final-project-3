package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-3/dto"
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/helpers"
	"github.com/hacktiv8-ks07-g04/final-project-3/repository/category_repository"
)

type CategoryService interface {
	CreateCategory(payload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.MessageErr)
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
		StatusCode: 200,
		Message:    "Successfully created new category",
		Data: dto.NewCategoryDataResponse{
			CategoryID: category.CategoryID,
			Type:       category.Type,
			CreatedAt:  category.CreatedAt,
		},
	}

	return &response, nil
}
