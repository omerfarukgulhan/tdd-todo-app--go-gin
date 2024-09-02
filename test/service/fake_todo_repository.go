package service

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
	"todo-app--go-gin/domain"
	"todo-app--go-gin/persistence"
)

type FakeTodoRepository struct {
	todos []domain.Todo
}

func NewFakeTodoRepository(initialTodos []domain.Todo) persistence.ITodoRepository {
	return &FakeTodoRepository{
		todos: initialTodos,
	}
}

func (fakeTodoRepository *FakeTodoRepository) GetAllTodos() ([]domain.Todo, error) {
	return fakeTodoRepository.todos, nil
}

func (fakeTodoRepository *FakeTodoRepository) GetTodoById(todoId int) (domain.Todo, error) {
	for _, todo := range fakeTodoRepository.todos {
		if todo.Id == todoId {
			return todo, nil
		}
	}

	return domain.Todo{}, errors.New(fmt.Sprintf("Todo with id %d not found", todoId))
}

func (fakeTodoRepository *FakeTodoRepository) GetAllTodosByUserId(userId int) ([]domain.Todo, error) {
	var userTodos []domain.Todo
	for _, todo := range fakeTodoRepository.todos {
		if todo.UserId == userId {
			userTodos = append(userTodos, todo)
		}
	}

	return userTodos, nil
}

func (fakeTodoRepository *FakeTodoRepository) AddTodo(todo domain.Todo) (domain.Todo, error) {
	todo.Id = len(fakeTodoRepository.todos) + 1
	fakeTodoRepository.todos = append(fakeTodoRepository.todos, todo)

	return todo, nil
}

func (fakeTodoRepository *FakeTodoRepository) UpdateTodo(todoId int, updatedTodo domain.Todo) (domain.Todo, error) {
	for i, todo := range fakeTodoRepository.todos {
		if todo.Id == todoId {
			updatedTodo.Id = todo.Id
			updatedTodo.UserId = todo.UserId
			updatedTodo.CreatedAt = todo.CreatedAt
			updatedTodo.UpdatedAt = time.Now()
			fakeTodoRepository.todos[i] = updatedTodo
			return fakeTodoRepository.todos[i], nil
		}
	}

	return domain.Todo{}, errors.New(fmt.Sprintf("Todo with id %d not found", todoId))
}

func (fakeTodoRepository *FakeTodoRepository) DeleteTodo(todoId int) error {
	for i, todo := range fakeTodoRepository.todos {
		if todo.Id == todoId {
			fakeTodoRepository.todos = append(fakeTodoRepository.todos[:i], fakeTodoRepository.todos[i+1:]...)
			return nil
		}
	}

	return errors.New(fmt.Sprintf("Todo with id %d not found", todoId))
}
