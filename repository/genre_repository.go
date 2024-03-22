package repository

import "final-project/models"

type GenreRepository interface {
	Save(genre models.Genre)
	Update(genre models.Genre)
	Delete(id int)
	FindById(id int) (models.Genre, error)
	FindAll() []models.Genre
}