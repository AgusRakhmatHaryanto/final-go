package repository

import (
	"final-project/helper"
	"final-project/models"

	"gorm.io/gorm"
)

type GenreRepositoryImpl struct {
	db *gorm.DB
}

func NewGenreRepositoryImpl(db *gorm.DB) GenreRepository {
	return &GenreRepositoryImpl{db: db}
}

// Delete implements GenreRepository.
func (g *GenreRepositoryImpl) Delete(id int) {
	var genre models.Genre

	result := g.db.Where("id = ?", id).Delete(&genre)
	helper.ErrorPanic(result.Error)
}

// FindAll implements GenreRepository.
func (g *GenreRepositoryImpl) FindAll() []models.Genre {
	var genres []models.Genre
	result := g.db.Find(&genres)
	helper.ErrorPanic(result.Error)
	return genres
}

// FindById implements GenreRepository.
func (g *GenreRepositoryImpl) FindById(id int) (models.Genre, error) {
	var genre models.Genre
	result := g.db.Where("id = ?", id).Find(&genre)
	if result.RowsAffected == 0 {
		return genre, result.Error
	}
	return genre, nil
}

// Save implements GenreRepository.
func (g *GenreRepositoryImpl) Save(genre models.Genre) {
	result := g.db.Create(&genre)
	helper.ErrorPanic(result.Error)
}

// Update implements GenreRepository.
func (g *GenreRepositoryImpl) Update(genre models.Genre) {
	var updateGenre = models.Genre{
		ID:   genre.ID,
		Name: genre.Name,
	}
	result := g.db.Model(&genre).Updates(updateGenre)
	helper.ErrorPanic(result.Error)
}

