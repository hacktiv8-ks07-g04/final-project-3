package category_repository

import (
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
)

type Repository interface {
	CreateCategory(category *entity.Category) errs.MessageErr
}
