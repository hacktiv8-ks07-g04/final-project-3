package category_repository

import (
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
)

type Repository interface {
	CreateCategory(category *entity.Category) errs.MessageErr
	DeleteCategory(id uint) errs.MessageErr
	UpdateCategory(id uint, category *entity.Category) (*entity.Category, errs.MessageErr)
	GetCategoryWithTask() ([]entity.Category, errs.MessageErr)
	GetCategoryById(id uint) (*entity.Category, errs.MessageErr)
}
