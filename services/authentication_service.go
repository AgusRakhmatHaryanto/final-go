package services

import (
	"final-project/data/request"
	"final-project/data/response"
)

type AuthenticationService interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.RegisterNewUserRequest)
	FindByEmail(email string) response.UserResponse
	FindByUsername(username string) response.UserResponse
}
