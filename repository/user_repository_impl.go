package repository

import (
	"errors"
	"final-project/data/request"
	"final-project/helper"
	"final-project/models"

	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	db *gorm.DB
}



func NewUserRepositoryImpl(db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{db: db}
}


// Delete implements UsersRepository.
func (u *UsersRepositoryImpl) Delete(userId int) {
	var users models.Users
	result := u.db.Where("id = ?", userId).Delete(&users)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UsersRepository.
func (u *UsersRepositoryImpl) FindAll() []models.Users {
	var users []models.Users
	result := u.db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

// FindById implements UsersRepository.
func (u *UsersRepositoryImpl) FindById(userId int) (models.Users, error) {
	var users models.Users
	result := u.db.Find(&users, userId)
	if result != nil {
		return users, nil
	} else {
		return users, errors.New("user not found")
	}
}

// Save implements UsersRepository.
func (u *UsersRepositoryImpl) Save(users models.Users) {
	result := u.db.Create(&users)
	helper.ErrorPanic(result.Error)
}

// Update implements UsersRepository.
func (u *UsersRepositoryImpl) Update(users models.Users) {
	var updateUser = request.UpdateUserRequest{
		ID:       users.ID,
		Username: users.Username,
		Email:    users.Email,
		Password: users.Password,
	}

	result := u.db.Model(&users).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}

// findByUsername implements UsersRepository.
func (u *UsersRepositoryImpl) FindByUsername(username string) (models.Users, error) {
	var users models.Users
	result := u.db.First(&users, "username = ?", username)

	if result.Error != nil {
		return users, errors.New("Invalid username")
	}

	return users, nil
}

// findByEmail implements UsersRepository.
func (u *UsersRepositoryImpl) FindByEmail(email string) (models.Users, error) {
	var users models.Users
	result := u.db.Where("email = ?", email).First(&users)
	if result.Error != nil {
		return users, errors.New("user not found")
	}
	return users, nil
}
