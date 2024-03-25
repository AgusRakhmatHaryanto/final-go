package services

import (
	"final-project/data/request"
	"final-project/data/response"
	"final-project/helper"
	"final-project/models"
	"final-project/repository"
	"log"

	"github.com/go-playground/validator/v10"
)

type movieServiceImpl struct {
	movieRepository repository.MovieRepository
	validator       *validator.Validate
	awardService    AwardService
	genreService    GenreService
}

func NewMovieServiceImpl(movieRepository repository.MovieRepository, validate *validator.Validate, awardService AwardService, genreService GenreService) MovieService {
	return &movieServiceImpl{
		movieRepository: movieRepository,
		validator:       validate,
		awardService:    awardService,
		genreService:    genreService,
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
	log.Println(movie.ID)
	award := m.awardService.FindAwardById(movie.AwardID)
	genre := m.genreService.FindGenreById(movie.GenreID)
	log.Println(award)
	log.Println(genre)
	log.Println(movie.ID)
	data_award := response.AwardResponse{
		ID:   award.ID,
		Year: award.Year,
		Title: award.Title,
	}

	data_genre:= response.GenreResponse{
		ID:   genre.ID,
		Name: genre.Name,
	}
	return response.MovieGetAllResponse{
		ID:      movie.ID,
		Title:   movie.Title,
		Year:    movie.Year,
		AwardID: movie.AwardID,
		GenreID: movie.GenreID,
		Award:   data_award,
		Genre:   data_genre,
	}
}

// GetAllMovies implements MovieService.
func (m *movieServiceImpl) GetAllMovies() []response.MovieGetAllResponse {
	result := m.movieRepository.FindAll()

	var movies []response.MovieGetAllResponse
	for _, movie := range result {
		award := m.awardService.FindAwardById(movie.AwardID)
		genre := m.genreService.FindGenreById(movie.GenreID)
		movies = append(movies, response.MovieGetAllResponse{
			ID:      movie.ID,
			Title:   movie.Title,
			Year:    movie.Year,
			AwardID: movie.AwardID,
			GenreID: movie.GenreID,
			Award:   award,
			Genre:   genre,
		})
	}

	return movies
}

// SaveMovie implements MovieService.
func (m *movieServiceImpl) SaveMovie(movie request.CreateNewMoviesRequest) {
	newMovie := models.Movie{
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
