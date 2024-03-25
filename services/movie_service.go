package services

import (
	"final-project/data/request"
	"final-project/data/response"
)

type MovieService interface {
	UpdateMovie(movie request.UpdateMoviesRequest)
	SaveMovie(movie request.CreateNewMoviesRequest)
	DeleteMovie(movieId int)
	FindMovieById(movieId int) response.MovieGetAllResponse
	GetAllMovies() []response.MovieGetAllResponse
	
}