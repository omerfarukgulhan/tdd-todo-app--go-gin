package service

import (
	"todo-app--go-gin/common/util/security"
	"todo-app--go-gin/domain"
	"todo-app--go-gin/domain/request"
	"todo-app--go-gin/domain/response"
	"todo-app--go-gin/persistence"
)

type IUserService interface {
	GetAllUsers() ([]response.UserResponse, error)
	//GetUserById(userId int) (response.UserResponse, error)
	//GetUserByEmail(userId int) ([]response.UserResponse, error)
	AddUser(userCreate request.UserCreate) (response.UserResponse, error)
	//UpdateUser(userId int, UserUpdate request.UserUpdate) (response.UserResponse, error)
	//DeleteUser(userId int) error
}

type UserService struct {
	userRepository persistence.IUserRepository
}

func NewUserService(userRepository persistence.IUserRepository) IUserService {
	return &UserService{userRepository: userRepository}
}

func (userService UserService) GetAllUsers() ([]response.UserResponse, error) {
	users, err := userService.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return convertUsersToResponses(users), nil
}

func (userService UserService) AddUser(userCreate request.UserCreate) (response.UserResponse, error) {
	hashedPassword, err := security.HashPassword(userCreate.Password)
	if err != nil {
		return response.UserResponse{}, err
	}

	user, err := userService.userRepository.AddUser(domain.User{
		Username: userCreate.Username,
		Email:    userCreate.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.NewUserResponse(user), nil
}

func convertUsersToResponses(users []domain.User) []response.UserResponse {
	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, response.NewUserResponse(user))
	}

	return userResponses
}
