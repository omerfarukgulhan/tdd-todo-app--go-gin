package service

import (
	"github.com/pkg/errors"
	"time"
	"todo-app--go-gin/domain"
	"todo-app--go-gin/domain/request"
	"todo-app--go-gin/domain/response"
	"todo-app--go-gin/persistence"
)

type ITodoService interface {
	GetAllTodos(userId int) ([]response.TodoResponse, error)
	GetTodoById(userId int, todoId int) (response.TodoResponse, error)
	AddTodo(todoCreate request.TodoCreate) (response.TodoResponse, error)
	UpdateTodo(todoId int, todoUpdate request.TodoUpdate) (response.TodoResponse, error)
	ToggleTodo(userId int, todoId int) (response.TodoResponse, error)
	DeleteTodo(userId int, todoId int) error
}

type TodoService struct {
	todoRepository persistence.ITodoRepository
}

func NewTodoService(todoRepository persistence.ITodoRepository) ITodoService {
	return &TodoService{todoRepository: todoRepository}
}

func (todoService TodoService) GetAllTodos(userId int) ([]response.TodoResponse, error) {
	todos, err := todoService.todoRepository.GetAllTodosByUserId(userId)
	if err != nil {
		return nil, err
	}

	return convertTodosToResponses(todos), nil
}

func (todoService TodoService) GetTodoById(userId int, todoId int) (response.TodoResponse, error) {
	todo, err := todoService.todoRepository.GetTodoById(todoId)
	if err != nil {
		return response.TodoResponse{}, err
	}

	if todo.UserId != userId {
		return response.TodoResponse{}, errors.New("This todo is not belongs to you")
	}

	return response.NewTodoResponse(todo), nil
}

func (todoService TodoService) AddTodo(todoCreate request.TodoCreate) (response.TodoResponse, error) {
	validationError := validateTodo(todoCreate)
	if validationError != nil {
		return response.TodoResponse{}, validationError
	}

	addedTodo, err := todoService.todoRepository.AddTodo(domain.Todo{
		UserId:      todoCreate.UserId,
		Title:       todoCreate.Title,
		Description: todoCreate.Description,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		return response.TodoResponse{}, errors.Wrap(err, "Failed to add new todo")
	}

	return response.NewTodoResponse(addedTodo), nil
}

func (todoService TodoService) UpdateTodo(todoId int, todoUpdate request.TodoUpdate) (response.TodoResponse, error) {
	validationError := validateTodo(todoUpdate)
	if validationError != nil {
		return response.TodoResponse{}, validationError
	}

	todo, err := todoService.todoRepository.GetTodoById(todoId)
	if err != nil {
		return response.TodoResponse{}, err
	}

	if todo.UserId != todoUpdate.UserId {
		return response.TodoResponse{}, errors.New("This todo is not belongs to you")
	}

	todo.UpdatedAt = time.Now()
	todo.Title = todoUpdate.Title
	todo.Description = todoUpdate.Description
	todo.IsCompleted = todoUpdate.IsCompleted

	_, err = todoService.todoRepository.UpdateTodo(todoId, todo)
	if err != nil {
		return response.TodoResponse{}, err
	}

	return response.NewTodoResponse(todo), nil
}

func (todoService TodoService) ToggleTodo(userId int, todoId int) (response.TodoResponse, error) {
	todo, err := todoService.todoRepository.GetTodoById(todoId)
	if err != nil {
		return response.TodoResponse{}, err
	}

	if todo.UserId != userId {
		return response.TodoResponse{}, errors.New("This todo is not belongs to you")
	}

	todo.IsCompleted = !todo.IsCompleted

	_, err = todoService.todoRepository.UpdateTodo(todoId, todo)
	if err != nil {
		return response.TodoResponse{}, err
	}

	return response.NewTodoResponse(todo), nil
}

func (todoService TodoService) DeleteTodo(userId int, todoId int) error {
	todo, err := todoService.GetTodoById(userId, todoId)
	if err != nil {
		return err
	}

	if todo.UserId != userId {
		return errors.New("This todo is not belongs to you")
	}

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

func convertTodosToResponses(todos []domain.Todo) []response.TodoResponse {
	var todoResponses []response.TodoResponse
	for _, todo := range todos {
		todoResponses = append(todoResponses, response.NewTodoResponse(todo))
	}

	return todoResponses
}
