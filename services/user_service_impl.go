package services

import (
	"final-project/data/request"
	"final-project/data/response"
	"final-project/helper"
	"final-project/models"
	"final-project/repository"
	"final-project/utils"
	"final-project/utils/enum"
	"final-project/validation"

	"github.com/go-playground/validator/v10"
)

type UsersServiceImpl struct {
	userRepository repository.UsersRepository
	valiadator     *validator.Validate
}

func NewUsersServiceImpl(usersRepository repository.UsersRepository, validate *validator.Validate) UsersService {
	return &UsersServiceImpl{
		userRepository: usersRepository,
		valiadator:     validate,
	}
}

// CreateUser implements UsersService.
func (u *UsersServiceImpl) CreateUser(users request.CreateNewUserRequest) {

	hashPassword, err := utils.HashPassword(users.Password)

	helper.ErrorPanic(err)

	role, errRole := validation.CheckEqual(users.Role, enum.RoleType)
	helper.ErrorPanic(errRole)
	newUser := models.Users{
		Username: users.Username,
		Email:    users.Email,
		Password: hashPassword,
		Role:      role,
	}

	u.userRepository.Save(newUser)

}

// DeleteUser implements UserService.
func (u *UsersServiceImpl) DeleteUser(UserId int) {
	u.userRepository.Delete(UserId)
}

// FindUserByEmail implements UserService.
func (u *UsersServiceImpl) FindUserByEmail(email string) response.UserResponse {
	userData, err := u.userRepository.FindByEmail(email)
	helper.ErrorPanic(err)

	user_res := response.UserResponse{
		ID:       userData.ID,
		Username: userData.Username,
		Email:    userData.Email,
		Password: userData.Password,
		Role:     userData.Role,
	}

	return user_res
}

// FindUserById implements UserService.
func (u *UsersServiceImpl) FindUserById(UserId int) response.UserResponse {
	userData, err := u.userRepository.FindById(UserId)
	helper.ErrorPanic(err)

	user_res := response.UserResponse{
		ID:       userData.ID,
		Username: userData.Username,
		Email:    userData.Email,
		Password: userData.Password,
	}

	return user_res
}

// FindUserByUsername implements UserService.
func (u *UsersServiceImpl) FindUserByUsername(username string) response.UserResponse {
	userData, err := u.userRepository.FindByUsername(username)
	helper.ErrorPanic(err)

	user_res := response.UserResponse{
		ID:       userData.ID,
		Username: userData.Username,
		Email:    userData.Email,
		Password: userData.Password,
		Role:     userData.Role,
	}

	return user_res
}

// GetAllUsers implements UserService.
func (u *UsersServiceImpl) GetAllUsers() []response.UserResponse {
	result := u.userRepository.FindAll()

	var user_res []response.UserResponse
	for _, user := range result {
		user_res = append(user_res, response.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
			Role:     user.Role,
		})

	}
	return user_res
}

// UpdateUser implements UserService.
func (u *UsersServiceImpl) UpdateUser(users request.UpdateUserRequest) {
	userData, err := u.userRepository.FindById(users.ID)
	helper.ErrorPanic(err)

	userData.Username = users.Username
	userData.Email = users.Email
	hashPassword, err := utils.HashPassword(users.Password)

	helper.ErrorPanic(err)
	userData.Password = hashPassword

	u.userRepository.Update(userData)
}
