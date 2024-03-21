package repository

import "final-project/models"

type UsersRepository interface {
	Save(users models.Users)
	Update(users models.Users)
	Delete(userId int)
	FindById(userId int) (models.Users, error)
	FindAll() []models.Users
	FindByUsername(username string) (models.Users, error)
	FindByEmail(email string) (models.Users, error)
}
