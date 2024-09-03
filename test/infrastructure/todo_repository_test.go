package infrastructure

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo-app--go-gin/domain"
)

func TestGetAllTodos(t *testing.T) {
	SetupData(ctx, dbPool)

	t.Run("GetAllTodos", func(t *testing.T) {
		actualTodos, _ := todoRepository.GetAllTodos()
		assert.Equal(t, 4, len(actualTodos))
	})

	ClearData(ctx, dbPool)
}

func TestGetTodoById(t *testing.T) {
	SetupData(ctx, dbPool)

	expectedTodo := domain.Todo{
		Id:          1,
		UserId:      1,
		Title:       "Buy groceries",
		Description: "Purchase fruits, vegetables, and bread",
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("GetTodoById", func(t *testing.T) {
		actualTodo, _ := todoRepository.GetTodoById(expectedTodo.Id)
		assert.Equal(t, expectedTodo.Title, actualTodo.Title)
		assert.Equal(t, expectedTodo.Description, actualTodo.Description)
	})

	ClearData(ctx, dbPool)
}

func TestGetAllTodosByUserId(t *testing.T) {
	SetupData(ctx, dbPool)

	expectedTodos := []domain.Todo{
		{
			Id:          1,
			UserId:      1,
			Title:       "Buy groceries",
			Description: "Purchase fruits, vegetables, and bread",
			IsCompleted: false,
			CreatedAt:   MustParseTime("2024-09-01T10:00:00"),
			UpdatedAt:   MustParseTime("2024-09-01T10:00:00"),
		},
		{
			Id:          2,
			UserId:      1,
			Title:       "Complete assignment",
			Description: "Finish the report for the upcoming meeting",
			IsCompleted: true,
			CreatedAt:   MustParseTime("2024-09-02T09:30:00"),
			UpdatedAt:   MustParseTime("2024-09-02T09:30:00"),
		},
		{
			Id:          3,
			UserId:      1,
			Title:       "Workout session",
			Description: "Attend the gym for a cardio session",
			IsCompleted: false,
			CreatedAt:   MustParseTime("2024-09-03T18:00:00"),
			UpdatedAt:   MustParseTime("2024-09-03T18:00:00"),
		},
	}

	t.Run("GetAllTodosByUserId", func(t *testing.T) {
		actualTodos, _ := todoRepository.GetAllTodosByUserId(expectedTodos[0].UserId)
		assert.Equal(t, len(expectedTodos), len(actualTodos))
		assert.Equal(t, expectedTodos[0].UserId, actualTodos[0].UserId)
	})

	ClearData(ctx, dbPool)
}

func TestAddTodo(t *testing.T) {
	SetupData(ctx, dbPool)

	newTodo := domain.Todo{
		Id:          5,
		UserId:      1,
		Title:       "Buy groceries",
		Description: "Purchase fruits, vegetables, and bread",
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("AddTodo", func(t *testing.T) {
		_, err := todoRepository.AddTodo(newTodo)
		fmt.Println(err)
		actualTodos, _ := todoRepository.GetAllTodos()
		assert.Equal(t, 5, len(actualTodos))
	})

	ClearData(ctx, dbPool)
}

func TestUpdateTodo(t *testing.T) {
	SetupData(ctx, dbPool)

	expectedTodo := domain.Todo{
		Id:          1,
		UserId:      1,
		Title:       "Buy groceries updated",
		Description: "Purchase fruits, vegetables, and bread",
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	updatedTodo := domain.Todo{
		UserId:      1,
		Title:       "Buy groceries updated",
		Description: "Purchase fruits, vegetables, and bread",
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("UpdateTodo", func(t *testing.T) {
		todoRepository.UpdateTodo(expectedTodo.Id, updatedTodo)
		actualTodos, _ := todoRepository.GetTodoById(expectedTodo.Id)
		assert.Equal(t, expectedTodo.Title, actualTodos.Title)
		assert.Equal(t, expectedTodo.Description, actualTodos.Description)
	})

	ClearData(ctx, dbPool)
}

func TestDeleteTodo(t *testing.T) {
	SetupData(ctx, dbPool)

	expectedTodos := []domain.Todo{
		{
			Id:          1,
			UserId:      1,
			Title:       "Buy groceries",
			Description: "Purchase fruits, vegetables, and bread",
			IsCompleted: false,
			CreatedAt:   MustParseTime("2024-09-01T10:00:00"),
			UpdatedAt:   MustParseTime("2024-09-01T10:00:00"),
		},
		{
			Id:          2,
			UserId:      1,
			Title:       "Complete assignment",
			Description: "Finish the report for the upcoming meeting",
			IsCompleted: true,
			CreatedAt:   MustParseTime("2024-09-02T09:30:00"),
			UpdatedAt:   MustParseTime("2024-09-02T09:30:00"),
		},
		{
			Id:          4,
			UserId:      2,
			Title:       "Read a book",
			Description: "Start reading a new novel",
			IsCompleted: true,
			CreatedAt:   MustParseTime("2024-09-04T20:00:00"),
			UpdatedAt:   MustParseTime("2024-09-04T20:00:00"),
		},
	}

	t.Run("DeleteTodo", func(t *testing.T) {
		todoRepository.DeleteTodo(3)
		actualTodos, _ := todoRepository.GetAllTodos()
		assert.Equal(t, len(expectedTodos), len(actualTodos))
	})

	ClearData(ctx, dbPool)
}
