package service

import (
	"github.com/pkg/errors"
	"regexp"
	"strings"
	"todo-app--go-gin/common/util/security"
	"todo-app--go-gin/domain"
	"todo-app--go-gin/domain/request"
	"todo-app--go-gin/domain/response"
	"todo-app--go-gin/persistence"
)

type IUserService interface {
	GetAllUsers() ([]response.UserResponse, error)
	//GetUserById(userId int) (response.UserResponse, error)
	GetUserByEmail(email string) (response.UserResponse, error)
	GetUserByEmailForValidation(email string) (domain.User, error)
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

func (userService UserService) GetUserByEmail(email string) (response.UserResponse, error) {
	user, err := userService.userRepository.GetUserByEmail(email)
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.NewUserResponse(user), nil
}

func (userService UserService) GetUserByEmailForValidation(email string) (domain.User, error) {
	user, err := userService.userRepository.GetUserByEmail(email)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (userService UserService) AddUser(userCreate request.UserCreate) (response.UserResponse, error) {
	validationError := validateUser(userCreate)
	if validationError != nil {
		return response.UserResponse{}, validationError
	}

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

func validateUser(userCreate request.UserCreate) error {
	if strings.TrimSpace(userCreate.Username) == "" {
		return errors.New("Username cannot be empty")
	}

	if !isValidEmail(userCreate.Email) {
		return errors.New("Invalid email format")
	}

	if len(userCreate.Password) < 5 {
		return errors.New("Password must be at least 8 characters long")
	}

	return nil
}

func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
