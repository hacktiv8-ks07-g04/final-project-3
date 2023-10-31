package helpers

import (
	"github.com/asaskevich/govalidator"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
)

func ValidateStruct(payload interface{}) errs.MessageErr {

	_, err := govalidator.ValidateStruct(payload)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}
