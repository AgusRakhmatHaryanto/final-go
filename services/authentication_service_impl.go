package services

import (
	"errors"
	"final-project/config"
	"final-project/data/request"
	"final-project/helper"
	"final-project/repository"
	"final-project/utils"
	"final-project/models"

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
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_user.ID, config.TokenSecret)
	helper.ErrorPanic(err_token)

	return token, nil

}

// Register implements AuthenticationService.
func (a *AuthenticationServiceImpl) Register(users request.CreateNewUserRequest)  {
	hashedPassword, err := utils.HashPassword(users.Password)
	
	helper.ErrorPanic(err)

	newUser:= models.Users{
		Username: users.Username,
		Email:    users.Email,
		Password: hashedPassword,
	}

	a.UsersRepository.Save(newUser)

}
