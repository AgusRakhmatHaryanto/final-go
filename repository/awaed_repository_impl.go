package repository

import (
	"final-project/helper"
	"final-project/models"

	"gorm.io/gorm"
)

type AwardRepositoryImpl struct {
	DB *gorm.DB
}

func NewAwardRepositoryImpl(db *gorm.DB) AwardRepository {
	return &AwardRepositoryImpl{DB: db}
}

// Delete implements AwardRepository.
func (a *AwardRepositoryImpl) Delete(id int) {
	var award models.Award

	result := a.DB.Where("id = ?", id).Delete(&award)
	helper.ErrorPanic(result.Error)
}

// FindAll implements AwardRepository.
func (a *AwardRepositoryImpl) FindAll() []models.Award {
	var awards []models.Award
	result := a.DB.Find(&awards)
	helper.ErrorPanic(result.Error)
	return awards
}

// FindAwardById implements AwardRepository.
func (a *AwardRepositoryImpl) FindById(id int) (models.Award, error) {
	var award models.Award

	result := a.DB.Where("id = ?", id).Find(&award)
	if result.RowsAffected == 0 {
		return award, result.Error
	}
	return award, nil
}


// Save implements AwardRepository.
func (a *AwardRepositoryImpl) Save(award models.Award) {
	result := a.DB.Create(&award)
	helper.ErrorPanic(result.Error)
}

// Update implements AwardRepository.
func (a *AwardRepositoryImpl) Update(award models.Award) {
	var updateAward = models.Award{
		ID:   award.ID,
		Title: award.Title,
		Year: award.Year,
	}
	result := a.DB.Model(&award).Updates(updateAward)
	helper.ErrorPanic(result.Error)
}
