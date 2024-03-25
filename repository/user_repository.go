package repository

import "final-project/models"

type UsersRepository interface {
	Save(users models.User)
	Update(users models.User)
	Delete(userId int)
	FindById(userId int) (models.User, error)
	FindAll() []models.User
	FindByUsername(username string) (models.User, error)
	FindByEmail(email string) (models.User, error)
}
