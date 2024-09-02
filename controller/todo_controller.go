package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"todo-app--go-gin/common/util/results"
	"todo-app--go-gin/controller/constants"
	"todo-app--go-gin/domain/request"
	"todo-app--go-gin/service"
)

type TodoController struct {
	todoService service.ITodoService
}

func NewTodoController(todoService service.ITodoService) *TodoController {
	return &TodoController{todoService: todoService}
}

func (todoController *TodoController) RegisterRoutes(router *gin.Engine) {
	todoGroup := router.Group("/todos")
	{
		todoGroup.GET("/", todoController.GetAllTodos)
		todoGroup.GET("/:id", todoController.GetTodoById)
		todoGroup.GET("/user/:userId", todoController.GetAllTodosByUserId)
		todoGroup.POST("/", todoController.AddTodo)
		todoGroup.PUT("/:id", todoController.UpdateTodo)
		todoGroup.PUT("/toggle/:id", todoController.ToggleTodo)
		todoGroup.DELETE("/:id", todoController.DeleteTodo)
	}
}

func (todoController *TodoController) GetAllTodos(ctx *gin.Context) {
	todos, err := todoController.todoService.GetAllTodos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, results.NewResult(false, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, results.NewDataResult(true, constants.DataFetched, todos))
}

func (todoController *TodoController) GetTodoById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, results.NewResult(false, "Invalid todo id"))
		return
	}

	todo, err := todoController.todoService.GetTodoById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, results.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, results.NewDataResult(true, constants.DataFetched, todo))
}

func (todoController *TodoController) GetAllTodosByUserId(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, results.NewResult(false, "Invalid user id"))
		return
	}

	todos, err := todoController.todoService.GetAllTodosByUserId(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, results.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, results.NewDataResult(true, constants.DataFetched, todos))
}

func (todoController *TodoController) AddTodo(ctx *gin.Context) {
	var newTodo request.TodoCreate
	if err := ctx.ShouldBindJSON(&newTodo); err != nil {
		log.Printf("error in bind")
		ctx.JSON(http.StatusBadRequest, results.NewResult(false, "Enter todo in valid format"))
		return
	}

	todo, err := todoController.todoService.AddTodo(newTodo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, results.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, results.NewDataResult(true, constants.DataAdded, todo))
}

func (todoController *TodoController) UpdateTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, results.NewResult(false, "Invalid todo id"))
		return
	}

	var updatedTodo request.TodoUpdate
	if err := ctx.ShouldBindJSON(&updatedTodo); err != nil {
		ctx.JSON(http.StatusBadRequest, results.NewResult(false, err.Error()))
		return
	}

	todo, err := todoController.todoService.UpdateTodo(id, updatedTodo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, results.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, results.NewDataResult(true, constants.DataUpdated, todo))
}

func (todoController *TodoController) ToggleTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, results.NewResult(false, "Invalid todo id"))
		return
	}

	todo, err := todoController.todoService.ToggleTodo(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, results.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, results.NewDataResult(true, constants.DataUpdated, todo))
}

func (todoController *TodoController) DeleteTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, results.NewResult(false, "Invalid todo id"))
		return
	}

	err = todoController.todoService.DeleteTodo(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, results.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, results.NewResult(true, constants.DataDeleted))
}
