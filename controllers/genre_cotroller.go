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

type GenreController struct {
	genreService services.GenreService
}

func NewGenreController(genreService services.GenreService) *GenreController {
	return &GenreController{
		genreService: genreService,
	}
}

func (c *GenreController) FindAllGenres(ctx *gin.Context) {
	log.Info().Msg("Get All Genres")
	genres := c.genreService.GetAllGenres()
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Get All Genres",
		Data:    genres,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *GenreController) FindGenreById(ctx *gin.Context) {
	log.Info().Msg("Get Genre By Id")
	genreId := ctx.Param("id")
	id, err := strconv.Atoi(genreId)
	helper.ErrorPanic(err)
	genre := c.genreService.FindGenreById(id)
	if genre.ID == 0 {
		webResponse := response.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT FOUND",
			Message: "Genre Not Found",
			Data:    nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Get Genre By Id",
		Data:    genre,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *GenreController) CreateGenre(ctx *gin.Context) {
	log.Info().Msg("Create Genre")
	genreRequest := request.CreateNewGenreRequest{}
	err := ctx.ShouldBind(&genreRequest)
	helper.ErrorPanic(err)
	c.genreService.SaveGenre(genreRequest)
	genre := response.CreateNewGenreResponse{
	
		Name: genreRequest.Name,
	}
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Create Genre",
		Data:    genre,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}


func (c *GenreController) UpdateGenre(ctx *gin.Context) {
	log.Info().Msg("Update Genre")
	genreRequest := request.UpdateGenreRequest{}
	err := ctx.ShouldBind(&genreRequest)
	helper.ErrorPanic(err)
	genreId := ctx.Param("id")
	id, err := strconv.Atoi(genreId)
	helper.ErrorPanic(err)
	genreRequest.ID = id
	
	genre := c.genreService.FindGenreById(id)
	if  genre.ID == 0 {
		webResponse := response.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT FOUND",
			Message: "Genre Not Found",
			Data:    nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}
	c.genreService.UpdateFenre(genreRequest)
	genre_res := response.UpdateGenreResponse{
		ID:   genreRequest.ID,
		Name: genreRequest.Name,
	}
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Update Genre",
		Data:    genre_res,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}


func (c *GenreController) DeleteGenre(ctx *gin.Context) {
	log.Info().Msg("Delete Genre")
	genreId := ctx.Param("id")
	id, err := strconv.Atoi(genreId)
	helper.ErrorPanic(err)
	c.genreService.DeleteGenre(id)
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Delete Genre",
		Data:    nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}