package services

import (
	"final-project/data/request"
	"final-project/data/response"
	"final-project/helper"
	"final-project/models"
	"final-project/repository"

	"github.com/go-playground/validator/v10"
)

type DirectorServiceImpl struct {
	directorRepository repository.DirectorRepository
	validator          *validator.Validate
}

func NewDirectorServiceImpl(directorRepository repository.DirectorRepository, validate *validator.Validate) DirectorService {
	return &DirectorServiceImpl{
		directorRepository: directorRepository,
		validator:          validate,
	}
}

// DeleteDirector implements DirectorService.
func (d *DirectorServiceImpl) DeleteDirector(directorId int) {
	d.directorRepository.Delete(directorId)
}

// FindAllDirector implements DirectorService.
func (d *DirectorServiceImpl) FindAllDirector() []response.DirectorGetAllResponse {
	director := d.directorRepository.FindAll()
	
	var director_res []response.DirectorGetAllResponse
	for _, v := range director {
		director_res = append(director_res, response.DirectorGetAllResponse{
			ID:      v.ID,
			Name:    v.Name,
			MovieID: v.MovieID,
		})
	}
	return director_res
}

// FindDirectorById implements DirectorService.
func (d *DirectorServiceImpl) FindDirectorById(directorId int) response.DirectorGetAllResponse {
	director, err := d.directorRepository.FindById(directorId)
	helper.ErrorPanic(err)
	return response.DirectorGetAllResponse{
		ID:      director.ID,
		Name:    director.Name,
		MovieID: director.MovieID,
	}
}

// SaveDirector implements DirectorService.
func (d *DirectorServiceImpl) SaveDirector(director request.CreateNewDirectorRequest) {
	newDirector := models.Director{
		Name:    director.Name,
		MovieID: director.MovieID,
	}

	d.directorRepository.Save(newDirector)
}

// UpdateDirector implements DirectorService.
func (d *DirectorServiceImpl) UpdateDirector(director request.UpdateDirectorRequest) {
	directorData, err := d.directorRepository.FindById(director.ID)
	helper.ErrorPanic(err)

	directorData.Name = director.Name
	directorData.MovieID = director.MovieID
	d.directorRepository.Update(directorData)
}

