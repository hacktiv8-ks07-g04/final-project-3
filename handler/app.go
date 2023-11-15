package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hacktiv8-ks07-g04/final-project-3/handler/category_handler"
	"github.com/hacktiv8-ks07-g04/final-project-3/handler/task_handler"
	"github.com/hacktiv8-ks07-g04/final-project-3/handler/user_handler"
	"github.com/hacktiv8-ks07-g04/final-project-3/infra/config"
	"github.com/hacktiv8-ks07-g04/final-project-3/infra/database"
	"github.com/hacktiv8-ks07-g04/final-project-3/repository/category_repository/category_pg"
	"github.com/hacktiv8-ks07-g04/final-project-3/repository/task_repository/task_pg"
	"github.com/hacktiv8-ks07-g04/final-project-3/repository/user_repository/user_pg"
	"github.com/hacktiv8-ks07-g04/final-project-3/service"
)

func StartApp() {
	config.LoadAppConfig()

	database.InitializedDatabase()

	var port = config.Server().Port

	db := database.GetDbInstance()

	userRepo := user_pg.UserInit(db)
	categoryRepo := category_pg.CategoryInit(db)
	taskRepo := task_pg.NewTaskPg(db)

	userService := service.NewUserService(userRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	taskService := service.NewTaskService(taskRepo)

	userHandler := user_handler.NewUserHandler(userService)
	categoryHandler := category_handler.NewCategoryHandler(categoryService)
	taskHandler := task_handler.NewTaskHandler(taskService)

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

	categoryRoute := r.Group("/categories")
	{
		categoryRoute.Use(authService.Authentication())

		categoryRoute.GET("", categoryHandler.GetCategoryWithTask)
		categoryRoute.POST("", authService.AdminAuthorization(), categoryHandler.CreateCategory)
		categoryRoute.PATCH("/:categoryId", authService.AdminAuthorization(), categoryHandler.UpdateCategory)
		categoryRoute.DELETE("/:categoryId", authService.AdminAuthorization(), categoryHandler.DeleteCategory)
	}

	taskRoute := r.Group("/tasks")
	{
		taskRoute.Use(authService.Authentication())

		taskRoute.POST("", taskHandler.CreateNewTask)
		taskRoute.GET("", taskHandler.GetTaskWithUser)
		taskRoute.PUT("/:taskId", taskHandler.UpdateTaskById)
	}

	r.Run(":" + port)
}
