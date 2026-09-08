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

	// ISSUE: PORT is read from env with no validation or default fallback. If unset,
	// `r.Run(":")` will be called with an empty string.
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

	// ISSUE: authService is wired directly against the user repository, bypassing the
	// UserService layer. Also, the auth middleware is defined in the service package
	// (auth_service.go), coupling the business layer to the HTTP framework. It should
	// live in a dedicated middleware package.
	authService := service.NewAuthService(userRepo)

	// ISSUE: gin.Default() has no CORS middleware configured. Cross-origin browser
	// requests will be blocked (or, without CORS, arbitrary origins are unrestricted)
	// — either way CORS should be explicitly configured.
	// ISSUE: No rate-limiting on the login endpoint, allowing brute-force attempts.
	r := gin.Default()
	userRoute := r.Group("/users")
	{

		userRoute.POST("/register", userHandler.RegisterNewUser)
		userRoute.POST("/login", userHandler.LoginUser)

		// ISSUE: auth middleware registered AFTER the register/login routes. In Gin,
		// middleware within a group only applies to routes registered after the
		// Use() call. It technically works but is fragile/confusing — a reader might
		// expect the middleware to protect all routes in the group. Move Use() to the
		// top for clarity.
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
		taskRoute.PATCH("/update-status/:taskId", taskHandler.UpdateTaskStatus)
		taskRoute.PATCH("/update-category/:taskId", taskHandler.UpdateTaskCategory)
		taskRoute.DELETE("/:taskId", taskHandler.DeleteTaskById)
	}

	// ISSUE: r.Run() is a blocking call with no graceful shutdown. When the app is
	// deployed and receives SIGTERM/SIGINT, the server is killed abruptly without
	// closing the DB connection or finishing in-flight requests. Should use
	// http.Server with server.Shutdown(ctx) on signal.
	// ISSUE: The error returned by r.Run() (e.g. port already in use) is ignored.
	r.Run(":" + port)
}
