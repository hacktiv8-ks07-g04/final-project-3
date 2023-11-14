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

// get category by id
func (c *categoryPG) GetCategoryById(id uint) (*entity.Category, errs.MessageErr) {
	var category entity.Category

	if err := c.db.First(&category, id).Error; err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &category, nil
}

// update category
func (c *categoryPG) UpdateCategory(id uint, category *entity.Category) (*entity.Category, errs.MessageErr) {
	var categoryData entity.Category

	if err := c.db.First(&categoryData, id).Error; err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	categoryData.Type = category.Type

	if err := c.db.Save(&categoryData).Error; err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &categoryData, nil
}

// delete category
func (c *categoryPG) DeleteCategory(id uint) errs.MessageErr {
	category, err := c.GetCategoryById(id)
	if err != nil {
		return err
	}

	if err := c.db.Delete(category).Error; err != nil {
		return errs.NewInternalServerError(err.Error())
	}

	return nil
}
