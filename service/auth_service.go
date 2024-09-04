package service

import (
	"todo-app--go-gin/common/util/security"
	"todo-app--go-gin/domain/request"
)

type IAuthService interface {
	Register(userCreate request.UserCreate) (string, error)
	Login(signInCredentials request.SignInCredentials) (string, error)
}

type AuthService struct {
	userService IUserService
}

func NewAuthService(userService IUserService) IAuthService {
	return &AuthService{userService: userService}
}

func (authService AuthService) Register(userCreate request.UserCreate) (string, error) {
	user, err := authService.userService.AddUser(userCreate)
	if err != nil {
		return "", err
	}

	token, err := security.GenerateToken(user.Id, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (authService AuthService) Login(signInCredentials request.SignInCredentials) (string, error) {
	user, err := authService.userService.GetUserByEmailForValidation(signInCredentials.Email)
	if err != nil {
		return "", err
	}

	if security.CheckPasswordHash(user.Password, signInCredentials.Password) {
		return "", err
	}

	token, err := security.GenerateToken(user.Id, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
