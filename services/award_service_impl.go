package services

import (
	"final-project/data/request"
	"final-project/data/response"
	"final-project/helper"
	"final-project/models"
	"final-project/repository"

	"github.com/go-playground/validator/v10"
)

type AwardServiceImpl struct {
	awardRepository repository.AwardRepository
	validator       *validator.Validate
}

func NewAwardServiceImpl(awardRepository repository.AwardRepository, validate *validator.Validate) AwardService {
	return &AwardServiceImpl{
		awardRepository: awardRepository,
		validator:       validate,
	}
}

// DeleteAward implements AwardService.
func (a *AwardServiceImpl) DeleteAward(awardId int) {
	a.awardRepository.Delete(awardId)
}

// FindAwardById implements AwardService.
func (a *AwardServiceImpl) FindAwardById(awardId int) response.AwardResponse {
	award, err := a.awardRepository.FindById(awardId)
	helper.ErrorPanic(err)

	award_res := response.AwardResponse{
		ID:   award.ID,
		Year: award.Year,
		Title: award.Title,
	}
	return award_res
}

// GetAllAwards implements AwardService.
func (a *AwardServiceImpl) GetAllAwards() []response.AwardResponse {
	result := a.awardRepository.FindAll()

	var awards []response.AwardResponse
	for _, award := range result {
		award_res := response.AwardResponse{
			ID:   award.ID,
			Year: award.Year,
			Title: award.Title,
		}
		awards = append(awards, award_res)
	}
	return awards
}

// SaveAward implements AwardService.
func (a *AwardServiceImpl) SaveAward(award request.CreateNewAwardRequest) {
	newAward := models.Award{
		Title: award.Title,
		Year:  award.Year,
	}

	a.awardRepository.Save(newAward)
}

// UpdateAward implements AwardService.
func (a *AwardServiceImpl) UpdateAward(award request.UpdateAwardRequest) {
	awardData, err := a.awardRepository.FindById(award.ID)
	helper.ErrorPanic(err)

	awardData.Title = award.Title
	awardData.Year = award.Year
	a.awardRepository.Update(awardData)
}

