package services

import (
	"final-project/data/request"
	"final-project/data/response"
	"final-project/helper"
	"final-project/models"
	"final-project/repository"

	"github.com/go-playground/validator/v10"
)

type movieServiceImpl struct {
	movieRepository repository.MovieRepository
	validator       *validator.Validate
}

func NewMovieServiceImpl(movieRepository repository.MovieRepository, validate *validator.Validate) MovieService {
	return &movieServiceImpl{
		movieRepository: movieRepository,
		validator:       validate,
	}
}

// DeleteMovie implements MovieService.
func (m *movieServiceImpl) DeleteMovie(movieId int) {
	m.movieRepository.Delete(movieId)
}

// FindMovieById implements MovieService.
func (m *movieServiceImpl) FindMovieById(movieId int) response.MovieGetAllResponse {
	movie, err := m.movieRepository.FindById(movieId)

	helper.ErrorPanic(err)

	return response.MovieGetAllResponse{
		ID:      movie.ID,
		Title:   movie.Title,
		Year:    movie.Year,
		AwardID: movie.AwardID,
		GenreID: movie.GenreID,
	}
}

// GetAllMovies implements MovieService.
func (m *movieServiceImpl) GetAllMovies() []response.MovieGetAllResponse {
	result  := m.movieRepository.FindAll()

	var movies []response.MovieGetAllResponse
	for _, movie := range result {
		movies = append(movies, response.MovieGetAllResponse{
			ID:      movie.ID,
			Title:   movie.Title,
			Year:    movie.Year,
			AwardID: movie.AwardID,
			GenreID: movie.GenreID,
		})
	}

	return movies
}

// SaveMovie implements MovieService.
func (m *movieServiceImpl) SaveMovie(movie request.CreateNewMoviesRequest) {
	newMovie := models.Movies{
		Title:   movie.Title,
		Year:    movie.Year,
		AwardID: movie.AwardID,
		GenreID: movie.GenreID,
	}

	m.movieRepository.Save(newMovie)
}

// UpdateMovie implements MovieService.
func (m *movieServiceImpl) UpdateMovie(movie request.UpdateMoviesRequest) {
	movieData, err := m.movieRepository.FindById(movie.ID)

	helper.ErrorPanic(err)

	movieData.Title = movie.Title
	movieData.Year = movie.Year
	movieData.AwardID = movie.AwardID
	movieData.GenreID = movie.GenreID
	m.movieRepository.Update(movieData)
}

