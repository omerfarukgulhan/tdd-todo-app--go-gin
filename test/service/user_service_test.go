package service

import (
	"github.com/go-playground/assert/v2"
	"testing"
	"todo-app--go-gin/domain/request"
)

func Test_ShouldGetAllUser(t *testing.T) {
	t.Run("ShouldGetAllUser", func(t *testing.T) {
		actualTodos, _ := userService.GetAllUsers()
		assert.Equal(t, 4, len(actualTodos))
	})
}

func Test_ShouldAddUser(t *testing.T) {
	t.Run("ShouldAddUser", func(t *testing.T) {
		userService.AddUser(request.UserCreate{
			Username: "user1",
			Email:    "user1@mail.com",
			Password: "12345",
		})
		actualUsers, _ := userService.GetAllUsers()
		assert.Equal(t, 5, len(actualUsers))
	})
}
