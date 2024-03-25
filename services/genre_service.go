package services

import (
	"final-project/data/request"
	"final-project/data/response"
)

type GenreService interface {
	UpdateGenre(genre request.UpdateGenreRequest)
	SaveGenre(genre request.CreateNewGenreRequest)
	DeleteGenre(genreId int)
	FindGenreById(genreId int) response.GenreResponse
	GetAllGenres() []response.GenreResponse
}