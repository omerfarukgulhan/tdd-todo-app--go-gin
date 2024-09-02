package response

import (
	"time"
	"todo-app--go-gin/domain"
)

type TodoResponse struct {
	UserId      int
	Title       string
	Description string
	IsCompleted bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTodoResponse(todo domain.Todo) TodoResponse {
	return TodoResponse{
		UserId:      todo.UserId,
		Title:       todo.Title,
		Description: todo.Description,
		IsCompleted: todo.IsCompleted,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}
