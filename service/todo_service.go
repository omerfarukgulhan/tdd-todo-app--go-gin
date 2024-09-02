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
	UpdateTodo(todoId int, todo domain.Todo) (domain.Todo, error)
	DeleteTodo(todoId int) error
}

type TodoService struct {
	todoRepository persistence.ITodoRepository
}

func NewTodoService(todoRepository persistence.ITodoRepository) ITodoService {
	return &TodoService{todoRepository: todoRepository}
}

func (todoService TodoService) GetAllTodos() ([]domain.Todo, error) {
	return todoService.GetAllTodos()
}

func (todoService TodoService) GetTodoById(todoId int) (domain.Todo, error) {
	return todoService.todoRepository.GetTodoById(todoId)
}

func (todoService TodoService) GetAllTodosByUserId(userId int) ([]domain.Todo, error) {
	return todoService.todoRepository.GetAllTodosByUserId(userId)
}

func (todoService TodoService) AddTodo(todoCreate request.TodoCreate) (domain.Todo, error) {
	validationError := validateTodoCreate(todoCreate)
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

func (todoService TodoService) UpdateTodo(todoId int, todo domain.Todo) (domain.Todo, error) {
	return todoService.UpdateTodo(todoId, todo)
}

func (todoService TodoService) DeleteTodo(todoId int) error {
	return todoService.todoRepository.DeleteTodo(todoId)
}

func validateTodoCreate(todoCreate request.TodoCreate) error {
	if len(todoCreate.Title) <= 3 {
		return errors.New("Todo title must be at least 3 character long")
	} else if len(todoCreate.Description) <= 5 {
		return errors.New("Todo description must be at least 5 character long")
	}

	return nil
}
