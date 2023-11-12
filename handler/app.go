package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hacktiv8-ks07-g04/final-project-3/infra/config"
	"github.com/hacktiv8-ks07-g04/final-project-3/infra/database"
	"github.com/hacktiv8-ks07-g04/final-project-3/repository/user_repository/user_pg"
	"github.com/hacktiv8-ks07-g04/final-project-3/service"
)

func StartApp() {
	config.LoadAppConfig()

	database.InitializedDatabase()

	var port = config.Server().Port

	db := database.GetDbInstance()

	userRepo := user_pg.UserInit(db)

	userService := service.NewUserService(userRepo)

	userHandler := NewUserHandler(userService)

	authService := service.NewAuthService(userRepo)

	r := gin.Default()
	userRoute := r.Group("/users")
	{

		userRoute.POST("/register", userHandler.RegisterNewUser)
		userRoute.POST("/login", userHandler.LoginUser)

		userRoute.Use(authService.Authentication())

		userRoute.PUT("/update-account", userHandler.UpdateUser)
		userRoute.DELETE("/delete-account", userHandler.DeleteUser)
	}

	r.Run(":" + port)
}
