package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-3/repository/user_repository"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
}

type authService struct {
	userRepo user_repository.Repository
}

func NewAuthService(userRepo user_repository.Repository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invalidTokenErr = errs.NewUnauthenticatedError("Invalid token")
		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		err := user.ValidateToken(bearerToken)

		if err != nil {
			fmt.Printf("[Authentication]: %s\n", err.Error())
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		result, err := a.userRepo.GetUserByEmail(user.Email)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		_ = result

		ctx.Set("userData", user)

		ctx.Next()
	}
}
