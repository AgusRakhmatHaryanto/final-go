package repository

import (
	"final-project/helper"
	"final-project/models"

	"gorm.io/gorm"
)

type MovieRepositoryImpl struct {
	DB *gorm.DB
}

func NewMovieRepositoryImpl(db *gorm.DB) MovieRepository {
	return &MovieRepositoryImpl{DB: db}
}

// FindAwardById implements MovieRepository.
func (a *MovieRepositoryImpl) FindAwardById(id int) (models.Award, error) {
	var award models.Award

	result := a.DB.Where("id = ?", id).Find(&award)
	helper.ErrorPanic(result.Error)
	return award, nil
}

// FindGenreById implements MovieRepository.
func (g *MovieRepositoryImpl) FindGenreById(id int) (models.Genre, error) {
	var genre models.Genre

	result := g.DB.Where("id = ?", id).Find(&genre)
	helper.ErrorPanic(result.Error)
	return genre, nil
}


// FindAllAwards implements MovieRepository.
func (a *MovieRepositoryImpl) FindAllAwards() []models.Award {
	var awards []models.Award

	result := a.DB.Find(&awards)
	helper.ErrorPanic(result.Error)
	return awards
}

// FindAllGenres implements MovieRepository.
func (g *MovieRepositoryImpl) FindAllGenres() []models.Genre {
	var genres []models.Genre
	result := g.DB.Find(&genres)
	helper.ErrorPanic(result.Error)
	return genres
}

// Delete implements MovieRepository.
func (m *MovieRepositoryImpl) Delete(id int) {
	var movie models.Movie

	result := m.DB.Where("id = ?", id).Delete(&movie)
	helper.ErrorPanic(result.Error)
}

// FindAll implements MovieRepository.
func (m *MovieRepositoryImpl) FindAll() []models.Movie {
	var movies []models.Movie
	result := m.DB.Find(&movies)
	helper.ErrorPanic(result.Error)
	return movies
}

// FindById implements MovieRepository.
func (m *MovieRepositoryImpl) FindById(id int) (models.Movie, error) {
	var movie models.Movie
	result := m.DB.Where("id = ?", id).Find(&movie)
	if result.RowsAffected == 0 {
		return movie, result.Error
	}
	return movie, nil
}

// Save implements MovieRepository.
func (m *MovieRepositoryImpl) Save(movie models.Movie) {
	result := m.DB.Create(&movie)
	helper.ErrorPanic(result.Error)
}

// Update implements MovieRepository.
func (m *MovieRepositoryImpl) Update(movie models.Movie) {
	var updateMovie = models.Movie{
		ID:      movie.ID,
		Title:   movie.Title,
		Year:    movie.Year,
		AwardID: movie.AwardID,
		GenreID: movie.GenreID,
	}
	result := m.DB.Model(&movie).Updates(updateMovie)
	helper.ErrorPanic(result.Error)
	if result.RowsAffected == 0 {
		helper.ErrorPanic(result.Error)
	}
}
