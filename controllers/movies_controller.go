package controllers

import (
	"final-project/data/request"
	"final-project/data/response"
	"final-project/helper"
	// "final-project/models"
	"final-project/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type MovieController struct {
	movieService services.MovieService
}

func NewMovieController(movieService services.MovieService) *MovieController {
	return &MovieController{
		movieService: movieService,
	}
}

func (c *MovieController) FindAllMovies(ctx *gin.Context) {
	log.Info().Msg("Get All Movies")
	movies := c.movieService.GetAllMovies()
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Get All Movies",
		Data:    movies,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (c *MovieController) FindMovieById(ctx *gin.Context) {
	log.Info().Msg("Get Movie By Id")
	movieId := ctx.Param("id")
	id, err := strconv.Atoi(movieId)
	helper.ErrorPanic(err)
	movie := c.movieService.FindMovieById(id)
	if movie.ID == 0 {
		webResponse := response.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT FOUND",
			Message: "Movie Not Found",
			Data:    nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Get Movie By Id",
		Data:    movie,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}


func (c *MovieController) CreateMovie(ctx *gin.Context) {
	log.Info().Msg("Create Movie")
	createMovieRequest:= request.CreateNewMoviesRequest{}
	err := ctx.ShouldBindJSON(&createMovieRequest)
	helper.ErrorPanic(err)
	c.movieService.SaveMovie(createMovieRequest)
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Create Movie",
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (c *MovieController) UpdateMovie(ctx *gin.Context) {
	log.Info().Msg("Update Movie")
	movieRequest := request.UpdateMoviesRequest{}
	err := ctx.ShouldBindJSON(&movieRequest)
	helper.ErrorPanic(err)
	movieId := ctx.Param("id")
	id, err := strconv.Atoi(movieId)
	helper.ErrorPanic(err)
	movieRequest.ID = id
	c.movieService.UpdateMovie(movieRequest)
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Update Movie",
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (c *MovieController) DeleteMovie(ctx *gin.Context) {
	log.Info().Msg("Delete Movie")
	movieId := ctx.Param("id")
	id, err := strconv.Atoi(movieId)
	helper.ErrorPanic(err)
	c.movieService.DeleteMovie(id)
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Delete Movie",
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}