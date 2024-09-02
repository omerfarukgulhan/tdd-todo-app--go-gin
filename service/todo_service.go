package service

import (
	"github.com/pkg/errors"
	"time"
	"todo-app--go-gin/domain"
	"todo-app--go-gin/domain/request"
	"todo-app--go-gin/persistence"
)

type ITodoService interface {
	GetAllTodos() ([]domain.Todo, error)
	GetTodoById(todoId int) (domain.Todo, error)
	GetAllTodosByUserId(userId int) ([]domain.Todo, error)
	AddTodo(todoCreate request.TodoCreate) (domain.Todo, error)
	UpdateTodo(todoId int, todoUpdate request.TodoUpdate) (domain.Todo, error)
	DeleteTodo(todoId int) error
}

type TodoService struct {
	todoRepository persistence.ITodoRepository
}

func NewTodoService(todoRepository persistence.ITodoRepository) ITodoService {
	return &TodoService{todoRepository: todoRepository}
}

func (todoService TodoService) GetAllTodos() ([]domain.Todo, error) {
	return todoService.todoRepository.GetAllTodos()
}

func (todoService TodoService) GetTodoById(todoId int) (domain.Todo, error) {
	return todoService.todoRepository.GetTodoById(todoId)
}

func (todoService TodoService) GetAllTodosByUserId(userId int) ([]domain.Todo, error) {
	return todoService.todoRepository.GetAllTodosByUserId(userId)
}

func (todoService TodoService) AddTodo(todoCreate request.TodoCreate) (domain.Todo, error) {
	validationError := validateTodo(todoCreate)
	if validationError != nil {
		return domain.Todo{}, validationError
	}

	return todoService.todoRepository.AddTodo(domain.Todo{
		UserId:      1,
		Title:       todoCreate.Title,
		Description: todoCreate.Description,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
}

func (todoService TodoService) UpdateTodo(todoId int, todoUpdate request.TodoUpdate) (domain.Todo, error) {
	validationError := validateTodo(todoUpdate)
	if validationError != nil {
		return domain.Todo{}, validationError
	}
	existingTodo, err := todoService.todoRepository.GetTodoById(todoId)
	if err != nil {
		return domain.Todo{}, err
	}

	updatedTodo := domain.Todo{
		Id:          existingTodo.Id,
		UserId:      existingTodo.UserId,
		Title:       todoUpdate.Title,
		Description: todoUpdate.Description,
		IsCompleted: todoUpdate.IsCompleted,
		CreatedAt:   existingTodo.CreatedAt,
		UpdatedAt:   time.Now(),
	}

	return todoService.todoRepository.UpdateTodo(todoId, updatedTodo)
}

func (todoService TodoService) DeleteTodo(todoId int) error {
	return todoService.todoRepository.DeleteTodo(todoId)
}

func validateTodo(todo interface{}) error {
	switch t := todo.(type) {
	case request.TodoCreate:
		if len(t.Title) <= 3 {
			return errors.New("Todo title must be at least 3 characters long")
		} else if len(t.Description) <= 5 {
			return errors.New("Todo description must be at least 5 characters long")
		}
	case request.TodoUpdate:
		if len(t.Title) <= 3 {
			return errors.New("Todo title must be at least 3 characters long")
		} else if len(t.Description) <= 5 {
			return errors.New("Todo description must be at least 5 characters long")
		}
	default:
		return errors.New("Unsupported type")
	}

	return nil
}
