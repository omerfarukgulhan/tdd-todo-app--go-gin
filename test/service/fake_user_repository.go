package service

import (
	"fmt"
	"github.com/pkg/errors"
	"todo-app--go-gin/domain"
	"todo-app--go-gin/persistence"
)

type FakeUserRepository struct {
	users []domain.User
}

func NewFakeUserRepository(initialUsers []domain.User) persistence.IUserRepository {
	return &FakeUserRepository{
		users: initialUsers,
	}
}

func (fakeUserRepository *FakeUserRepository) GetAllUsers() ([]domain.User, error) {
	return fakeUserRepository.users, nil
}

func (fakeUserRepository *FakeUserRepository) GetUserById(userId int) (domain.User, error) {
	for _, user := range fakeUserRepository.users {
		if userId == user.Id {
			return user, nil
		}
	}

	return domain.User{}, errors.New(fmt.Sprintf("User with id %d not found", userId))
}

func (fakeUserRepository *FakeUserRepository) GetUserByEmail(email string) (domain.User, error) {
	for _, user := range fakeUserRepository.users {
		if email == user.Email {
			return user, nil
		}
	}

	return domain.User{}, errors.New(fmt.Sprintf("User with email %s not found", email))
}

func (fakeUserRepository *FakeUserRepository) AddUser(user domain.User) (domain.User, error) {
	user.Id = len(fakeUserRepository.users) + 1
	fakeUserRepository.users = append(fakeUserRepository.users, user)

	return user, nil
}

func (fakeUserRepository *FakeUserRepository) UpdateUser(userId int, userToUpdate domain.User) (domain.User, error) {
	for i, user := range fakeUserRepository.users {
		if user.Id == userId {
			fakeUserRepository.users[i] = userToUpdate
			return user, nil
		}
	}

	return domain.User{}, errors.New(fmt.Sprintf("User with id %d not found", userId))
}

func (fakeUserRepository *FakeUserRepository) DeleteUser(userId int) error {
	for i, user := range fakeUserRepository.users {
		if user.Id == userId {
			fakeUserRepository.users = append(fakeUserRepository.users[:i], fakeUserRepository.users[i+1:]...)
			return nil
		}
	}

	return errors.New(fmt.Sprintf("Todo with id %d not found", userId))
}
