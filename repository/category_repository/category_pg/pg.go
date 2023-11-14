package category_pg

import (
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-3/repository/category_repository"
	"gorm.io/gorm"
)

type categoryPG struct {
	db *gorm.DB
}

func CategoryInit(db *gorm.DB) category_repository.Repository {
	return &categoryPG{
		db: db,
	}
}

// create category
func (c *categoryPG) CreateCategory(category *entity.Category) errs.MessageErr {
	if err := c.db.Create(&category).Error; err != nil {
		return errs.NewInternalServerError(err.Error())
	}
	return nil
}

// get category with all associated task
func (c *categoryPG) GetCategoryWithTask() ([]entity.Category, errs.MessageErr) {
	var categories []entity.Category
	if err := c.db.Preload("Task").Find(&categories).Error; err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	return categories, nil
}
