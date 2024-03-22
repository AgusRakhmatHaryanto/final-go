package repository

import "final-project/models"

type MovieRepository interface {
	Save(movie models.Movies)
	Update(movie models.Movies)
	Delete(id int)
	FindById(id int) (models.Movies, error)
	FindAll() []models.Movies
}