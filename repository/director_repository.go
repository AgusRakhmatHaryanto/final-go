package repository

import "final-project/models"

type DirectorRepository interface {
	Save(director models.Director)
	Update(director models.Director)
	Delete(id int)
	FindById(id int) (models.Director, error)
	FindAll() []models.Director
}