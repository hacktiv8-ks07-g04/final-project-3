package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-3/repository/user_repository"
)

// ISSUE: Middleware is defined in the service package. Middleware is an HTTP-framework
// concern (depends on gin.HandlerFunc, ctx.Set, ctx.MustGet) and should live in a
// dedicated middleware/ package instead of the service (business logic) layer.
type AuthService interface {
	Authentication() gin.HandlerFunc
	AdminAuthorization() gin.HandlerFunc
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
			// ISSUE: fmt.Printf used instead of structured logging.
			// Debug-style print left in production path.
			fmt.Printf("[Authentication]: %s\n", err.Error())
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		result, err := a.userRepo.GetUserByEmail(user.Email)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		// ISSUE: The fetched `result` is discarded. The user data stored in context
		// comes from JWT claims only, not from the DB. If a user is deleted after
		// their token was issued, the token is still accepted (stale data). Also adds
		// an unnecessary DB round-trip on every authenticated request.
		_ = result

		ctx.Set("userData", user)

		// ISSUE: DEAD CODE — at this point `err` is ALWAYS nil (the code already
		// returned on both earlier error checks). If err were somehow non-nil here,
		// this block would neither call Next() nor abort, causing the request to hang.
		if err == nil {
			ctx.Next()
		}
	}
}

// Admin Authorization
func (a *authService) AdminAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ISSUE: ctx.MustGet panics if the key is missing, and the unchecked type
		// assertion .(entity.User) panics if the stored value is not of that type.
		// Should use ctx.Get() with an ok check to avoid crashing the server.
		user := ctx.MustGet("userData").(entity.User)

		// ISSUE: Role is not embedded in the JWT, so this does a DB query on every
		// admin request to check the role. Adding role to the token would avoid this.
		acc, err := a.userRepo.GetUserById(user.ID)
		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if acc.Role != "admin" {
			err := errs.NewUnauthorizedError("Unauthorized to edit this data")
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		ctx.Next()
	}
}
