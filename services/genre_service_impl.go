package services

import (
	"final-project/data/request"
	"final-project/data/response"
	"final-project/helper"
	"final-project/models"
	"final-project/repository"

	"github.com/go-playground/validator/v10"
)

type GenreServiceImpl struct {
	genreRepository repository.GenreRepository
	validator       *validator.Validate
}

func NewGenreServiceImpl(genreRepository repository.GenreRepository, validate *validator.Validate) GenreService {
	return &GenreServiceImpl{
		genreRepository: genreRepository,
		validator:       validate,
	}
}

// DeleteGenre implements GenreService.
func (g *GenreServiceImpl) DeleteGenre(genreId int) {
	g.genreRepository.Delete(genreId)
}

// FindGenreById implements GenreService.
func (g *GenreServiceImpl) FindGenreById(genreId int) response.GenreResponse {
	genre, err := g.genreRepository.FindById(genreId)
	helper.ErrorPanic(err)

	genre_res := response.GenreResponse{
		ID:   genre.ID,
		Name: genre.Name,
	}

	return genre_res
	
}

// GetAllGenres implements GenreService.
func (g *GenreServiceImpl) GetAllGenres() []response.GenreResponse {
	genres := g.genreRepository.FindAll()
	var genre_res []response.GenreResponse
	for _, genre := range genres {
		genre_res = append(genre_res, response.GenreResponse{
			ID:   genre.ID,
			Name: genre.Name,
		})
	}
	return genre_res
}

// SaveGenre implements GenreService.
func (g *GenreServiceImpl) SaveGenre(genre request.CreateNewGenreRequest) {
	g.genreRepository.Save(models.Genre{
		Name: genre.Name,
	})
}

// UpdateFenre implements GenreService.
func (g *GenreServiceImpl) UpdateFenre(genre request.UpdateGenreRequest) {
	genreData, err := g.genreRepository.FindById(genre.ID)
	helper.ErrorPanic(err)

	genreData.Name = genre.Name
	g.genreRepository.Update(genreData)
}

