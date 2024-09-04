package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"todo-app--go-gin/domain"
)

func TestGetAllUsers(t *testing.T) {
	SetupData(ctx, dbPool)

	expectedUsers := []domain.User{
		{
			Id:       1,
			Username: "user1",
			Email:    "user1@mail.com",
			Password: "12345",
		},
		{
			Id:       2,
			Username: "user2",
			Email:    "user2@mail.com",
			Password: "12345",
		},
		{
			Id:       3,
			Username: "user3",
			Email:    "user3@mail.com",
			Password: "12345",
		},
		{
			Id:       4,
			Username: "user4",
			Email:    "user4@mail.com",
			Password: "12345",
		},
	}
	t.Run("GetAllUsers", func(t *testing.T) {
		actualUsers, _ := userRepository.GetAllUsers()
		assert.Equal(t, expectedUsers, actualUsers)
	})

	ClearData(ctx, dbPool)
}

func TestGetUserById(t *testing.T) {
	SetupData(ctx, dbPool)

	expectedUser := domain.User{
		Id:       3,
		Username: "user3",
		Email:    "user3@mail.com",
		Password: "12345",
	}

	t.Run("GetUserById", func(t *testing.T) {
		actualUser, _ := userRepository.GetUserById(3)
		assert.Equal(t, expectedUser, actualUser)
	})

	ClearData(ctx, dbPool)
}

func TestGetUserByEmail(t *testing.T) {
	SetupData(ctx, dbPool)

	expectedUser := domain.User{
		Id:       3,
		Username: "user3",
		Email:    "user3@mail.com",
		Password: "12345",
	}

	t.Run("GetUserByEmail", func(t *testing.T) {
		actualUser, _ := userRepository.GetUserByEmail("user3@mail.com")
		assert.Equal(t, expectedUser, actualUser)
	})

	ClearData(ctx, dbPool)
}

func TestAddUser(t *testing.T) {
	expectedUsers := []domain.User{
		{
			Id:       1,
			Username: "user1",
			Email:    "user1@mail.com",
			Password: "12345",
		},
	}

	t.Run("AddUser", func(t *testing.T) {
		userRepository.AddUser(domain.User{
			Username: "user1",
			Email:    "user1@mail.com",
			Password: "12345",
		})
		actualUsers, _ := userRepository.GetAllUsers()
		assert.Equal(t, expectedUsers, actualUsers)
	})

	ClearData(ctx, dbPool)
}

func TestUpdateUser(t *testing.T) {
	SetupData(ctx, dbPool)

	expectedUser := domain.User{
		Id:       1,
		Username: "user1 updated",
		Email:    "user1@mail.com",
		Password: "123456789",
	}

	t.Run("UpdateUser", func(t *testing.T) {
		userRepository.UpdateUser(1, domain.User{
			Username: "user1 updated",
			Email:    "user1@mail.com",
			Password: "123456789",
		})
		actualUser, _ := userRepository.GetUserById(1)
		assert.Equal(t, expectedUser, actualUser)
	})

	ClearData(ctx, dbPool)
}

func TestDeleteUser(t *testing.T) {
	SetupData(ctx, dbPool)

	t.Run("DeleteUser", func(t *testing.T) {
		userRepository.DeleteUser(1)
		actualUsers, _ := userRepository.GetAllUsers()
		assert.Equal(t, 3, len(actualUsers))
	})

	ClearData(ctx, dbPool)
}
