package response

import (
	"todo-app--go-gin/domain"
)

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewUserResponse(user domain.User) UserResponse {
	return UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}
