package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"todo-app--go-gin/domain"
	"todo-app--go-gin/domain/request"
	"todo-app--go-gin/domain/response"
)

func Test_ShouldGetAllTodo(t *testing.T) {
	t.Run("ShouldGetAllTodo", func(t *testing.T) {
		actualTodos, _ := todoService.GetAllTodos()
		assert.Equal(t, 4, len(actualTodos))
	})
}

func Test_ShouldGetTodoById(t *testing.T) {
	expectedTodo := response.NewTodoResponse(domain.Todo{
		Id:          1,
		UserId:      1,
		Title:       "Buy groceries",
		Description: "Purchase fruits, vegetables, and bread",
		IsCompleted: false,
		CreatedAt:   mustParseTime("2024-09-01T10:00:00"),
		UpdatedAt:   mustParseTime("2024-09-01T10:00:00"),
	})

	t.Run("ShouldGetTodoById", func(t *testing.T) {
		actualTodo, _ := todoService.GetTodoById(1)
		assert.Equal(t, expectedTodo, actualTodo)
	})
}

func Test_ShouldNotGetTodoById(t *testing.T) {
	expectedTodo := domain.Todo{
		Id:          1,
		UserId:      1,
		Title:       "Buy groceries",
		Description: "Purchase fruits, vegetables, and bread",
		IsCompleted: false,
		CreatedAt:   mustParseTime("2024-09-01T10:00:00"),
		UpdatedAt:   mustParseTime("2024-09-01T10:00:00"),
	}

	t.Run("ShouldNotGetTodoById", func(t *testing.T) {
		actualTodo, _ := todoService.GetTodoById(2)
		assert.NotEqual(t, expectedTodo, actualTodo)
	})
}

func Test_ShouldNotGetTodoByIdInvalidId(t *testing.T) {
	t.Run("ShouldNotGetTodoByIdInvalidId", func(t *testing.T) {
		_, err := todoService.GetTodoById(5)
		assert.Equal(t, "Todo with id 5 not found", err.Error())
	})
}

func Test_ShouldGetTodosByUserId(t *testing.T) {
	t.Run("ShouldGetTodosByUserId", func(t *testing.T) {
		userTodos, _ := todoService.GetAllTodosByUserId(1)
		assert.Equal(t, 3, len(userTodos))
	})
}

func Test_ShouldAddTodo(t *testing.T) {
	t.Run("ShouldAddTodo", func(t *testing.T) {
		todoService.AddTodo(request.TodoCreate{Title: "title", Description: "description"})
		actualTodos, _ := todoService.GetAllTodos()
		assert.Equal(t, 5, len(actualTodos))
	})
}

func Test_WhenTodoTitleLengthLessThan3(t *testing.T) {
	t.Run("WhenTodoTitleLengthLessThan3", func(t *testing.T) {
		_, err := todoService.AddTodo(request.TodoCreate{Title: "tt", Description: "description"})
		actualTodos, _ := todoService.GetAllTodos()
		assert.Equal(t, 4, len(actualTodos))
		assert.Equal(t, "Todo title must be at least 3 characters long", err.Error())
	})
}

func Test_WhenTodoDescriptionLengthLessThan5(t *testing.T) {
	t.Run("WhenTodoDescriptionLengthLessThan5", func(t *testing.T) {
		_, err := todoService.AddTodo(request.TodoCreate{Title: "title", Description: "dsc"})
		actualTodos, _ := todoService.GetAllTodos()
		assert.Equal(t, 4, len(actualTodos))
		assert.Equal(t, "Todo description must be at least 5 characters long", err.Error())
	})
}

func Test_ShouldUpdateTodo(t *testing.T) {
	exceptedTodo := domain.Todo{
		Id:          1,
		UserId:      1,
		Title:       "Buy groceries updated",
		Description: "Purchase fruits, vegetables, and bread",
		IsCompleted: true,
		CreatedAt:   mustParseTime("2024-09-01T10:00:00"),
		UpdatedAt:   mustParseTime("2024-09-01T10:00:00"),
	}

	t.Run("ShouldUpdateTodo", func(t *testing.T) {
		todoService.UpdateTodo(1, request.TodoUpdate{
			Title:       "Buy groceries updated",
			Description: "Purchase fruits, vegetables, and bread",
			IsCompleted: true,
		})
		actualTodo, _ := todoService.GetTodoById(1)
		assert.Equal(t, exceptedTodo.Title, actualTodo.Title)
		assert.Equal(t, exceptedTodo.IsCompleted, actualTodo.IsCompleted)
	})
}

func Test_ShouldNotUpdateTodoInvalidId(t *testing.T) {
	t.Run("ShouldNotUpdateTodoInvalidId", func(t *testing.T) {
		_, err := todoService.UpdateTodo(6, request.TodoUpdate{
			Title:       "Buy groceries updated",
			Description: "Purchase fruits, vegetables, and bread",
			IsCompleted: true,
		})
		_, actualErr := todoService.GetTodoById(6)
		assert.Equal(t, err.Error(), actualErr.Error())
	})
}

func Test_ShouldDeleteTodo(t *testing.T) {
	t.Run("ShouldDeleteTodo", func(t *testing.T) {
		todoService.DeleteTodo(1)
		actualTodos, _ := todoService.GetAllTodos()
		assert.Equal(t, 3, len(actualTodos))
	})
}

func Test_ShouldNotDeleteTodoInvalidId(t *testing.T) {
	t.Run("ShouldNotDeleteTodoInvalidId", func(t *testing.T) {
		err := todoService.DeleteTodo(6)
		_, actualErr := todoService.GetTodoById(6)
		assert.Equal(t, err.Error(), actualErr.Error())
	})
}
