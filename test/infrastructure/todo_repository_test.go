package infrastructure

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
	"todo-app--go-gin/common/postgresql"
	"todo-app--go-gin/domain"
	"todo-app--go-gin/persistence"
)

var todoRepository persistence.ITodoRepository
var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()

	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "5432",
		UserName:              "postgres",
		Password:              "153515",
		DbName:                "workshops",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "10s",
	})

	todoRepository = persistence.NewTodoRepository(dbPool)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func setupData(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbPool)
}

func clearData(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)
}

func parseTime(timeStr string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05", timeStr)
}

func mustParseTime(timeStr string) time.Time {
	t, err := parseTime(timeStr)
	if err != nil {
		panic(err)
	}
	return t
}

func TestGetAllTodos(t *testing.T) {
	setupData(ctx, dbPool)

	expectedTodos := []domain.Todo{
		{
			Id:          1,
			UserId:      1,
			Title:       "Buy groceries",
			Description: "Purchase fruits, vegetables, and bread",
			IsCompleted: false,
			CreatedAt:   mustParseTime("2024-09-01T10:00:00"),
			UpdatedAt:   mustParseTime("2024-09-01T10:00:00"),
		},
		{
			Id:          2,
			UserId:      1,
			Title:       "Complete assignment",
			Description: "Finish the report for the upcoming meeting",
			IsCompleted: true,
			CreatedAt:   mustParseTime("2024-09-02T09:30:00"),
			UpdatedAt:   mustParseTime("2024-09-02T09:30:00"),
		},
		{
			Id:          3,
			UserId:      1,
			Title:       "Workout session",
			Description: "Attend the gym for a cardio session",
			IsCompleted: false,
			CreatedAt:   mustParseTime("2024-09-03T18:00:00"),
			UpdatedAt:   mustParseTime("2024-09-03T18:00:00"),
		},
		{
			Id:          4,
			UserId:      2,
			Title:       "Read a book",
			Description: "Start reading a new novel",
			IsCompleted: true,
			CreatedAt:   mustParseTime("2024-09-04T20:00:00"),
			UpdatedAt:   mustParseTime("2024-09-04T20:00:00"),
		},
	}

	t.Run("GetAllTodos", func(t *testing.T) {
		actualTodos, _ := todoRepository.GetAllTodos()
		assert.Equal(t, len(expectedTodos), len(actualTodos))
	})

	clearData(ctx, dbPool)
}

func TestGetTodoById(t *testing.T) {
	setupData(ctx, dbPool)

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

	clearData(ctx, dbPool)
}

func TestGetAllTodosByUserId(t *testing.T) {
	setupData(ctx, dbPool)

	expectedTodos := []domain.Todo{
		{
			Id:          1,
			UserId:      1,
			Title:       "Buy groceries",
			Description: "Purchase fruits, vegetables, and bread",
			IsCompleted: false,
			CreatedAt:   mustParseTime("2024-09-01T10:00:00"),
			UpdatedAt:   mustParseTime("2024-09-01T10:00:00"),
		},
		{
			Id:          2,
			UserId:      1,
			Title:       "Complete assignment",
			Description: "Finish the report for the upcoming meeting",
			IsCompleted: true,
			CreatedAt:   mustParseTime("2024-09-02T09:30:00"),
			UpdatedAt:   mustParseTime("2024-09-02T09:30:00"),
		},
		{
			Id:          3,
			UserId:      1,
			Title:       "Workout session",
			Description: "Attend the gym for a cardio session",
			IsCompleted: false,
			CreatedAt:   mustParseTime("2024-09-03T18:00:00"),
			UpdatedAt:   mustParseTime("2024-09-03T18:00:00"),
		},
	}

	t.Run("GetAllTodosByUserId", func(t *testing.T) {
		actualTodos, _ := todoRepository.GetAllTodosByUserId(expectedTodos[0].UserId)
		assert.Equal(t, len(expectedTodos), len(actualTodos))
		assert.Equal(t, expectedTodos[0].UserId, actualTodos[0].UserId)
	})

	clearData(ctx, dbPool)
}

func TestAddTodo(t *testing.T) {
	expectedTodos := []domain.Todo{
		{
			Id:          1,
			UserId:      1,
			Title:       "Buy groceries",
			Description: "Purchase fruits, vegetables, and bread",
			IsCompleted: false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	newTodo := domain.Todo{
		UserId:      1,
		Title:       "Buy groceries",
		Description: "Purchase fruits, vegetables, and bread",
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("AddTodo", func(t *testing.T) {
		todoRepository.AddTodo(newTodo)
		actualTodos, _ := todoRepository.GetAllTodos()
		assert.Equal(t, expectedTodos[0].Title, actualTodos[0].Title)
		assert.Equal(t, expectedTodos[0].Description, actualTodos[0].Description)
	})

	clearData(ctx, dbPool)
}

func TestUpdateTodo(t *testing.T) {
	setupData(ctx, dbPool)

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

	clearData(ctx, dbPool)
}

func TestDeleteTodo(t *testing.T) {
	setupData(ctx, dbPool)

	expectedTodos := []domain.Todo{
		{
			Id:          1,
			UserId:      1,
			Title:       "Buy groceries",
			Description: "Purchase fruits, vegetables, and bread",
			IsCompleted: false,
			CreatedAt:   mustParseTime("2024-09-01T10:00:00"),
			UpdatedAt:   mustParseTime("2024-09-01T10:00:00"),
		},
		{
			Id:          2,
			UserId:      1,
			Title:       "Complete assignment",
			Description: "Finish the report for the upcoming meeting",
			IsCompleted: true,
			CreatedAt:   mustParseTime("2024-09-02T09:30:00"),
			UpdatedAt:   mustParseTime("2024-09-02T09:30:00"),
		},
		{
			Id:          4,
			UserId:      2,
			Title:       "Read a book",
			Description: "Start reading a new novel",
			IsCompleted: true,
			CreatedAt:   mustParseTime("2024-09-04T20:00:00"),
			UpdatedAt:   mustParseTime("2024-09-04T20:00:00"),
		},
	}

	t.Run("DeleteTodo", func(t *testing.T) {
		todoRepository.DeleteTodo(3)
		actualTodos, _ := todoRepository.GetAllTodos()
		assert.Equal(t, len(expectedTodos), len(actualTodos))
	})

	clearData(ctx, dbPool)
}

//	func TestSetupData(t *testing.T) {
//		setupData(ctx, dbPool)
//	}
func TestClearData(t *testing.T) {
	clearData(ctx, dbPool)
}
