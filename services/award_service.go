package services

import (
	"final-project/data/request"
	"final-project/data/response"
)

type AwardService interface {
	UpdateAward(award request.UpdateAwardRequest)
	SaveAward(award request.CreateNewAwardRequest)
	DeleteAward(awardId int)
	FindAwardById(awardId int) response.AwardResponse
	GetAllAwards() []response.AwardResponse
}
