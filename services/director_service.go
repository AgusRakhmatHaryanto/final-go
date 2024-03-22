package services

import (
	"final-project/data/request"
	"final-project/data/response"
)

type DirectorService interface {
	UpdateDirector(director request.UpdateDirectorRequest)
	SaveDirector(director request.CreateNewDirectorRequest)
	DeleteDirector(directorId int)
	FindDirectorById(directorId int) response.DirectorGetAllResponse
	FindAllDirector() []response.DirectorGetAllResponse
}