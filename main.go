package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"todo-app--go-gin/common/app"
	"todo-app--go-gin/common/postgresql"
	"todo-app--go-gin/controller"
	"todo-app--go-gin/persistence"
	"todo-app--go-gin/service"
)

func main() {
	ctx := context.Background()
	server := gin.Default()
	configurationManager := app.NewConfigurationManager()
	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgreSqlConfig)

	todoRepo := persistence.NewTodoRepository(dbPool)
	todoService := service.NewTodoService(todoRepo)
	todoController := controller.NewTodoController(todoService)

	todoController.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
