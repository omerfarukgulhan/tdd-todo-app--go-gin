package service

import (
	"github.com/go-playground/assert/v2"
	"testing"
	"todo-app--go-gin/domain"
	"todo-app--go-gin/domain/request"
	"todo-app--go-gin/domain/response"
)

func Test_ShouldGetAllUser(t *testing.T) {
	t.Run("ShouldGetAllUser", func(t *testing.T) {
		actualUsers, _ := userService.GetAllUsers()
		assert.Equal(t, 4, len(actualUsers))
	})
}

func Test_ShouldGetUserById(t *testing.T) {
	t.Run("ShouldGetUserById", func(t *testing.T) {
		actualUser, _ := userService.GetUserById(1)
		expectedUser := response.UserResponse{
			Id:       1,
			Username: "user1",
			Email:    "user1@mail.com",
		}
		assert.Equal(t, expectedUser, (actualUser))
	})
}

func Test_ShouldNotGetUserById(t *testing.T) {
	t.Run("ShouldNotGetUserById", func(t *testing.T) {
		actualUser, err := userService.GetUserById(5)
		expectedUser := response.UserResponse{}
		assert.Equal(t, expectedUser, actualUser)
		assert.Equal(t, "User with id 5 not found", err.Error())
	})
}

func Test_ShouldGetUserByEmail(t *testing.T) {
	t.Run("ShouldGetUserByEmail", func(t *testing.T) {
		actualUser, _ := userService.GetUserByEmail("user1@mail.com")
		expectedUser := response.UserResponse{
			Id:       1,
			Username: "user1",
			Email:    "user1@mail.com",
		}
		assert.Equal(t, expectedUser, (actualUser))
	})
}

func Test_ShouldNotGetUserByEmail(t *testing.T) {
	t.Run("ShouldNotGetUserByEmail", func(t *testing.T) {
		actualUser, err := userService.GetUserByEmail("user123@mail.com")
		expectedUser := response.UserResponse{}
		assert.Equal(t, expectedUser, actualUser)
		assert.Equal(t, "User with email user123@mail.com not found", err.Error())
	})
}

func Test_ShouldGetUserByEmailForValidation(t *testing.T) {
	t.Run("ShouldGetUserByEmailForValidation", func(t *testing.T) {
		actualUser, _ := userService.GetUserByEmailForValidation("user1@mail.com")
		expectedUser := domain.User{
			Id:       1,
			Username: "user1",
			Email:    "user1@mail.com",
			Password: "$2a$10$xNG1.Ig/Q7kR5l4PDpMHcOw9/xEd.SJWdo2woUBBSM2MIKFoU9eTe",
		}
		assert.Equal(t, expectedUser, (actualUser))
	})
}

func Test_ShouldNotGetUserByEmailForValidation(t *testing.T) {
	t.Run("ShouldNotGetUserByEmailForValidation", func(t *testing.T) {
		actualUser, err := userService.GetUserByEmailForValidation("user123@mail.com")
		expectedUser := domain.User{}
		assert.Equal(t, expectedUser, actualUser)
		assert.Equal(t, "User with email user123@mail.com not found", err.Error())
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

func Test_ShouldNotAddUserValidationErrorUsername(t *testing.T) {
	t.Run("ShouldNotAddUserValidationErrorUsername", func(t *testing.T) {
		_, err := userService.AddUser(request.UserCreate{
			Username: "",
			Email:    "user1@mail.com",
			Password: "12345",
		})
		actualUsers, _ := userService.GetAllUsers()
		assert.Equal(t, 4, len(actualUsers))
		assert.Equal(t, "Username cannot be empty", err.Error())
	})
}

func Test_ShouldNotAddUserValidationErrorEmail(t *testing.T) {
	t.Run("ShouldNotAddUserValidationErrorEmail", func(t *testing.T) {
		_, err := userService.AddUser(request.UserCreate{
			Username: "user1",
			Email:    "user1mail.com",
			Password: "12345",
		})
		actualUsers, _ := userService.GetAllUsers()
		assert.Equal(t, 4, len(actualUsers))
		assert.Equal(t, "Invalid email format", err.Error())
	})
}

func Test_ShouldNotAddUserValidationErrorPassword(t *testing.T) {
	t.Run("ShouldNotAddUserValidationErrorPassword", func(t *testing.T) {
		_, err := userService.AddUser(request.UserCreate{
			Username: "user1",
			Email:    "user1@mail.com",
			Password: "1234",
		})
		actualUsers, _ := userService.GetAllUsers()
		assert.Equal(t, 4, len(actualUsers))
		assert.Equal(t, "Password must be at least 5 characters long", err.Error())
	})
}

func Test_ShouldNotAddUserErrorEmailAlreadyInUser(t *testing.T) {
	t.Run("ShouldAddUser", func(t *testing.T) {
		userService.AddUser(request.UserCreate{
			Username: "user5",
			Email:    "user5@mail.com",
			Password: "12345",
		})
		_, err := userService.AddUser(request.UserCreate{
			Username: "user5",
			Email:    "user5@mail.com",
			Password: "12345",
		})
		actualUsers, _ := userService.GetAllUsers()
		assert.Equal(t, 5, len(actualUsers))
		assert.Equal(t, "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)", err.Error())
	})
}

func Test_ShouldUpdateUser(t *testing.T) {
	expectedUser := response.UserResponse{
		Id:       1,
		Username: "user 1 updated",
		Email:    "user1@mail.com",
	}
	t.Run("ShouldUpdateUser", func(t *testing.T) {
		actualUser, _ := userService.UpdateUser(1, request.UserUpdate{Username: "user 1 updated"})
		assert.Equal(t, expectedUser, actualUser)
	})
}

func Test_ShouldNotUpdateUser(t *testing.T) {
	t.Run("ShouldNotUpdateUser", func(t *testing.T) {
		_, err := userService.UpdateUser(1, request.UserUpdate{})
		assert.Equal(t, "Username cannot be empty", err.Error())
	})
}

func Test_ShouldDeleteUser(t *testing.T) {
	t.Run("ShouldDeleteUser", func(t *testing.T) {
		userService.DeleteUser(1)
		actualUsers, _ := userService.GetAllUsers()
		assert.Equal(t, 3, len(actualUsers))
	})
}

func Test_ShouldNotDeleteUser(t *testing.T) {
	t.Run("ShouldNotDeleteUser", func(t *testing.T) {
		err := userService.DeleteUser(5)
		actualUsers, _ := userService.GetAllUsers()
		assert.Equal(t, 4, len(actualUsers))
		assert.Equal(t, "Todo with id 5 not found", err.Error())
	})
}
