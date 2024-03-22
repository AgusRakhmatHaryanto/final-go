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

// Delete implements MovieRepository.
func (m *MovieRepositoryImpl) Delete(id int) {
	var movie models.Movies

	result := m.DB.Where("id = ?", id).Delete(&movie)
	helper.ErrorPanic(result.Error)
}

// FindAll implements MovieRepository.
func (m *MovieRepositoryImpl) FindAll() []models.Movies {
	var movies []models.Movies
	result := m.DB.Find(&movies)
	helper.ErrorPanic(result.Error)
	return movies
}

// FindById implements MovieRepository.
func (m *MovieRepositoryImpl) FindById(id int) (models.Movies, error) {
	var movie models.Movies
	result := m.DB.Where("id = ?", id).Find(&movie)
	if result.RowsAffected == 0 {
		return movie, result.Error
	}
	return movie, nil
}

// Save implements MovieRepository.
func (m *MovieRepositoryImpl) Save(movie models.Movies) {
	result := m.DB.Create(&movie)
	helper.ErrorPanic(result.Error)
}

// Update implements MovieRepository.
func (m *MovieRepositoryImpl) Update(movie models.Movies) {
	var updateMovie = models.Movies{
		ID:   movie.ID,
		Title: movie.Title,
		Year: movie.Year,
		AwardID: movie.AwardID,
		GenreID: movie.GenreID,
	}
	result := m.DB.Model(&movie).Updates(updateMovie)
	helper.ErrorPanic(result.Error)
	if result.RowsAffected == 0 {
		helper.ErrorPanic(result.Error)
	}
}

