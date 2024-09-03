package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"time"
	"todo-app--go-gin/domain"
)

type ITodoRepository interface {
	GetAllTodos() ([]domain.Todo, error)
	GetTodoById(todoId int) (domain.Todo, error)
	GetAllTodosByUserId(userId int) ([]domain.Todo, error)
	AddTodo(todo domain.Todo) (domain.Todo, error)
	UpdateTodo(todoId int, todo domain.Todo) (domain.Todo, error)
	DeleteTodo(todoId int) error
}

type TodoRepository struct {
	dbPool *pgxpool.Pool
}

func NewTodoRepository(dbPool *pgxpool.Pool) ITodoRepository {
	return &TodoRepository{dbPool: dbPool}
}

func (todoRepository *TodoRepository) GetAllTodos() ([]domain.Todo, error) {
	ctx := context.Background()
	queryRow, err := todoRepository.dbPool.Query(ctx, "SELECT * FROM todos")
	if err != nil {
		return []domain.Todo{}, err
	}

	return extractTodosFromRows(queryRow), nil
}

func (todoRepository *TodoRepository) GetTodoById(todoId int) (domain.Todo, error) {
	ctx := context.Background()
	getByIdSql := `SELECT * FROM todos WHERE id = $1`
	queryRow := todoRepository.dbPool.QueryRow(ctx, getByIdSql, todoId)

	var id int
	var userId int
	var title string
	var description string
	var isCompleted bool
	var createdAt time.Time
	var updatedAt time.Time

	scanErr := queryRow.Scan(&id, &userId, &title, &description, &isCompleted, &createdAt, &updatedAt)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return domain.Todo{}, errors.New(fmt.Sprintf("Todo with id %d not found", todoId))
		}
		return domain.Todo{}, errors.New(fmt.Sprintf("Error while getting todo with id %d: %v", todoId, scanErr))
	}

	return domain.Todo{
		Id:          id,
		UserId:      userId,
		Title:       title,
		Description: description,
		IsCompleted: isCompleted,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}, nil
}

func (todoRepository *TodoRepository) GetAllTodosByUserId(userId int) ([]domain.Todo, error) {
	ctx := context.Background()
	getByIdSql := `SELECT * FROM todos WHERE user_id = $1`
	queryRow, err := todoRepository.dbPool.Query(ctx, getByIdSql, userId)
	if err != nil {
		return []domain.Todo{}, err
	}

	return extractTodosFromRows(queryRow), nil
}

func (todoRepository *TodoRepository) AddTodo(todo domain.Todo) (domain.Todo, error) {
	ctx := context.Background()
	insertSql := `INSERT INTO todos (user_id, title, description, is_completed, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var id int
	queryRow := todoRepository.dbPool.QueryRow(ctx, insertSql, todo.UserId, todo.Title, todo.Description, todo.IsCompleted, todo.CreatedAt, todo.UpdatedAt)
	scanErr := queryRow.Scan(&id)
	if scanErr != nil {
		return domain.Todo{}, scanErr
	}

	todo.Id = id

	return todo, nil
}

func (todoRepository *TodoRepository) UpdateTodo(todoId int, todo domain.Todo) (domain.Todo, error) {
	ctx := context.Background()
	updateTodoSql := `UPDATE todos SET user_id = $1, title = $2, description = $3, is_completed = $4, updated_at = $5 WHERE id = $6 RETURNING id, user_id, title, description, is_completed, updated_at;`
	queryRow := todoRepository.dbPool.QueryRow(ctx, updateTodoSql, todo.UserId, todo.Title, todo.Description, todo.IsCompleted, todo.UpdatedAt, todoId)
	scanErr := queryRow.Scan(&todo.Id, &todo.UserId, &todo.Title, &todo.Description, &todo.IsCompleted, &todo.UpdatedAt)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return domain.Todo{}, errors.New(fmt.Sprintf("Todo with id %d not found", todoId))
		}
		return domain.Todo{}, errors.New(fmt.Sprintf("Failed to update todo: %v", scanErr))
	}

	return todo, nil
}

func (todoRepository *TodoRepository) DeleteTodo(todoId int) error {
	ctx := context.Background()
	_, getErr := todoRepository.GetTodoById(todoId)
	if getErr != nil {
		return errors.New(fmt.Sprintf("Todo with id %d not found", todoId))
	}

	deleteSql := `DELETE FROM todos WHERE id = $1`
	_, err := todoRepository.dbPool.Exec(ctx, deleteSql, todoId)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while deleting todo with id %d", todoId))
	}

	return nil
}

func extractTodosFromRows(queryRow pgx.Rows) []domain.Todo {
	var todos = []domain.Todo{}
	for queryRow.Next() {
		var todo domain.Todo
		err := queryRow.Scan(
			&todo.Id,
			&todo.UserId,
			&todo.Title,
			&todo.Description,
			&todo.IsCompleted,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	if err := queryRow.Err(); err != nil {
		return todos
	}

	return todos
}
