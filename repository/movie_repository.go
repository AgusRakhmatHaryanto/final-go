package repository

import "final-project/models"

type MovieRepository interface {
	Save(movie models.Movie)
	Update(movie models.Movie)
	Delete(id int)
	FindById(id int) (models.Movie, error)
	FindAll() []models.Movie
	FindAllAwards() []models.Award
	FindAllGenres() []models.Genre
	FindAwardById(id int) (models.Award, error)
	FindGenreById(id int) (models.Genre, error)

}