package repository

import (
	"final-project/helper"
	"final-project/models"

	"gorm.io/gorm"
)

type DirectorRepositoryImpl struct {
	DB *gorm.DB
}

func NewDirectorRepositoryImpl(db *gorm.DB) DirectorRepository {
	return &DirectorRepositoryImpl{DB: db}
}

// Delete implements DirectorRepository.
func (d *DirectorRepositoryImpl) Delete(id int) {
	var director models.Director

	result := d.DB.Where("id = ?", id).Delete(&director)
	helper.ErrorPanic(result.Error)
}

// FindAll implements DirectorRepository.
func (d *DirectorRepositoryImpl) FindAll() []models.Director {
	var directors []models.Director

	result := d.DB.Find(&directors)
	helper.ErrorPanic(result.Error)
	return directors
}

// FindById implements DirectorRepository.
func (d *DirectorRepositoryImpl) FindById(id int) (models.Director, error) {
	var director models.Director
	result := d.DB.Where("id = ?", id).Find(&director)
	if result.RowsAffected == 0 {
		return director, result.Error
	}
	return director, nil
}

// Save implements DirectorRepository.
func (d *DirectorRepositoryImpl) Save(director models.Director) {
	result := d.DB.Create(&director)
	helper.ErrorPanic(result.Error)
}

// Update implements DirectorRepository.
func (d *DirectorRepositoryImpl) Update(director models.Director) {
	var updateDirector = models.Director{
		ID:   director.ID,
		Name: director.Name,
		MovieID: director.MovieID,
	}
	result := d.DB.Model(&director).Updates(updateDirector)
	helper.ErrorPanic(result.Error)
}

