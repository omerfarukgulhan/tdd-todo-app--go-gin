package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"todo-app--go-gin/common/app"
	"todo-app--go-gin/common/postgresql"
	"todo-app--go-gin/persistence"
	"todo-app--go-gin/service"
)

type MainRouter struct {
	authController *AuthController
	todoController *TodoController
}

func NewRouter(authController *AuthController, todoController *TodoController) *MainRouter {
	return &MainRouter{
		authController: authController,
		todoController: todoController,
	}
}

func (mainRouter *MainRouter) RegisterRoutes(server *gin.Engine) {
	mainRouter.authController.RegisterAuthRoutes(server)
	mainRouter.todoController.RegisterTodoRoutes(server)
}

func InitializeRouter() *gin.Engine {
	ctx := context.Background()
	server := gin.Default()

	configurationManager := app.NewConfigurationManager()
	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgreSqlConfig)

	todoRepo := persistence.NewTodoRepository(dbPool)
	todoService := service.NewTodoService(todoRepo)
	todoController := NewTodoController(todoService)

	userRepo := persistence.NewUserRepository(dbPool)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userService)
	authController := NewAuthController(authService)

	mainRouter := NewRouter(authController, todoController)
	mainRouter.RegisterRoutes(server)

	return server
}
