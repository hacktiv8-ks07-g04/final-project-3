package helpers

import (
	"github.com/asaskevich/govalidator"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
)

// ISSUE: Uses govalidator for validation, but Gin already vendors
// go-playground/validator/v10 (an indirect dependency in go.mod). Two validation
// libraries is redundant; also govalidator returns a raw error which is flattened
// into a single BadRequest here, losing per-field validation messages.
func ValidateStruct(payload interface{}) errs.MessageErr {

	_, err := govalidator.ValidateStruct(payload)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}
