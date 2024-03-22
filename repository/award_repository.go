package repository

import "final-project/models"

type AwardRepository interface {
	Save(award models.Award)
	Update(award models.Award)
	Delete(id int)
	FindById(id int) (models.Award, error)
	FindAll() []models.Award
}