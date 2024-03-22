package services

import (
	"final-project/data/request"
	"final-project/data/response"
)

type UsersService interface {
	UpdateUser(users request.UpdateUserRequest)
	DeleteUser(UserId int)
	FindUserById(UserId int) response.UserResponse
	FindUserByEmail(email string) response.UserResponse
	FindUserByUsername(username string) response.UserResponse
	GetAllUsers() []response.UserResponse
}
