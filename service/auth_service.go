package service

import (
	"todo-app--go-gin/domain/request"
)

type IAuthService interface {
	Register(userCreate request.UserCreate) error
	Login(signInCredentials request.SignInCredentials)
}

type AuthService struct {
	userService IUserService
}

func NewAuthService(userService IUserService) IAuthService {
	return &AuthService{userService: userService}
}

func (authService AuthService) Register(userCreate request.UserCreate) error {
	_, err := authService.userService.AddUser(userCreate)
	if err != nil {
		return err
	}

	return nil
}

func (authService AuthService) Login(signInCredentials request.SignInCredentials) {
	//TODO implement me
	panic("implement me")
}
