package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
)

func GetParamId(c *gin.Context, key string) (uint, errs.MessageErr) {
	value := c.Param(key)

	id, err := strconv.ParseUint(value, 10, 64)

	if err != nil {
		return 0, errs.NewBadRequest("invalid parameter id")
	}

	return uint(id), nil
}
