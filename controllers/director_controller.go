package controllers

import (
	"final-project/data/request"
	"final-project/data/response"
	"final-project/helper"
	"final-project/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type DirectorsController struct {
	directorService services.DirectorService
}

func NewDirectorController(directorService services.DirectorService) *DirectorsController {
	return &DirectorsController{
		directorService: directorService,
	}
}

func (c *DirectorsController) FindAllDirectors(ctx *gin.Context) {
	log.Info().Msg("Get All Directors")
	directors := c.directorService.FindAllDirector()
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Get All Director",
		Data:    directors,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *DirectorsController) FindDirectorById(ctx *gin.Context) {
	log.Info().Msg("Get Director By Id")
	directorId := ctx.Param("id")
	id, err := strconv.Atoi(directorId)
	helper.ErrorPanic(err)
	director := c.directorService.FindDirectorById(id)
	if director.ID == 0 {
		webResponse := response.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT FOUND",
			Message: "Director Not Found",
			Data:    nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Get Director By Id",
		Data:    director,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}


func (c *DirectorsController) CreateDirector(ctx *gin.Context) {
	log.Info().Msg("Create Director")
	directorRequest :=request.CreateNewDirectorRequest{}
	err := ctx.ShouldBindJSON(&directorRequest)
	helper.ErrorPanic(err)
	c.directorService.SaveDirector(directorRequest)

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Create Director",
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *DirectorsController) UpdateDirector(ctx *gin.Context) {
	log.Info().Msg("Update Director")
	directorRequest :=request.UpdateDirectorRequest{}
	err := ctx.ShouldBindJSON(&directorRequest)
	helper.ErrorPanic(err)
	directorId := ctx.Param("id")
	id, err := strconv.Atoi(directorId)
	helper.ErrorPanic(err)
	directorRequest.ID = id
	c.directorService.UpdateDirector(directorRequest)
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Update Director",
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}


func (c *DirectorsController) DeleteDirector(ctx *gin.Context) {
	log.Info().Msg("Delete Director")
	directorId := ctx.Param("id")
	id, err := strconv.Atoi(directorId)
	helper.ErrorPanic(err)
	c.directorService.DeleteDirector(id)
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Delete Director",
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}