package services

import (
	"errors"
	"final-project/config"
	"final-project/data/request"
	"final-project/data/response"
	"final-project/helper"
	"final-project/models"
	"final-project/repository"
	"final-project/utils"

	"github.com/go-playground/validator/v10"
)

type AuthenticationServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

func NewAuthenticationServiceImpl(usersRepository repository.UsersRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UsersRepository: usersRepository,
		Validate:        validate,
	}
}

// FIndByUsername implements AuthenticationService.
func (a *AuthenticationServiceImpl) FindByUsername(username string) response.UserResponse {
	user_data, err := a.UsersRepository.FindByUsername(username)
	helper.ErrorPanic(err)

	user_res := response.UserResponse{
		ID:       user_data.ID,
		Username: user_data.Username,
		Email:    user_data.Email,
		Password: user_data.Password,
		Role:     user_data.Role,
	}
	return user_res
}

// FindByEmail implements AuthenticationService.
func (a *AuthenticationServiceImpl) FindByEmail(email string) response.UserResponse {
	user_data, err := a.UsersRepository.FindByEmail(email)

	helper.ErrorPanic(err)

	user_res := response.UserResponse{
		ID:       user_data.ID,
		Username: user_data.Username,
		Email:    user_data.Email,
		Password: user_data.Password,
		Role:     user_data.Role,
	}
	return user_res
}

// Login implements AuthenticationService.
func (a *AuthenticationServiceImpl) Login(users request.LoginRequest) (string, error) {
	//find user in database
	new_user, err_user := a.UsersRepository.FindByEmail(users.Email)

	if err_user != nil {
		return "", errors.New("user not found")
	}

	config, _ := config.LoadConfig(".")

	verify_err := utils.VerifyPassword(new_user.Password, users.Password)

	if verify_err != nil {
		return "", errors.New("incorrect password")
	}

	//generate token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_user.ID, new_user.Role, config.TokenSecret)
	helper.ErrorPanic(err_token)

	return token, nil

}

// Register implements AuthenticationService.
func (a *AuthenticationServiceImpl) Register(users request.RegisterNewUserRequest) {
	hashed_password, err := utils.HashPassword(users.Password)

	helper.ErrorPanic(err)

	newUser := models.Users{
		Username: users.Username,
		Email:    users.Email,
		Password: hashed_password,
		Role:     users.Role,
	}

	a.UsersRepository.Save(newUser)

}
