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


type AwardController struct {
	awardService services.AwardService
}

func NewAwardController(awardService services.AwardService) *AwardController {
	return &AwardController{
		awardService: awardService,
	}
}

func (c AwardController) FindAllAwards(ctx *gin.Context) {
	log.Info().Msg("Get All Awards")
	awards := c.awardService.GetAllAwards()
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Get All Awards",
		Data:    awards,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}


func (c AwardController) FindAwardById(ctx *gin.Context) {
	log.Info().Msg("Get Award By Id")
	awardId := ctx.Param("id")
	id, err := strconv.Atoi(awardId)
	helper.ErrorPanic(err)
	award := c.awardService.FindAwardById(id)
	if award.ID == 0 {
		webResponse := response.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT FOUND",
			Message: "Award Not Found",
			Data:    nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Get Award By Id",
		Data:    award,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}


func (c AwardController) CreateAward(ctx *gin.Context) {
	log.Info().Msg("Create Award")
	awardRequest := request.CreateNewAwardRequest{}
	err := ctx.ShouldBind(&awardRequest)
	helper.ErrorPanic(err)
	c.awardService.SaveAward(awardRequest)
	award := response.AwardResponse{
		Title: awardRequest.Title,
		Year:  awardRequest.Year,
	}
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Create Award",
		Data:    award,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}


func (c AwardController) UpdateAward(ctx *gin.Context) {
	log.Info().Msg("Update Award")
	awardRequest := request.UpdateAwardRequest{}
	err := ctx.ShouldBind(&awardRequest)
	helper.ErrorPanic(err)
	awardId := ctx.Param("id")
	id, err := strconv.Atoi(awardId)
	helper.ErrorPanic(err)
	awardRequest.ID = id

	award := c.awardService.FindAwardById(id)
	if award.ID == 0 {
		webResponse := response.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT FOUND",
			Message: "Award Not Found",
			Data:    nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	c.awardService.UpdateAward(awardRequest)
	award_res := response.AwardResponse{
		ID:   awardRequest.ID,
		Year: awardRequest.Year,
		Title: awardRequest.Title,
	}
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Update Award",
		Data:    award_res,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}


func (c AwardController) DeleteAward(ctx *gin.Context) {
	log.Info().Msg("Delete Award")
	awardId := ctx.Param("id")
	id, err := strconv.Atoi(awardId)
	helper.ErrorPanic(err)
	award := c.awardService.FindAwardById(id)
	if award.ID == 0 {
		webResponse := response.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT FOUND",
			Message: "Award Not Found",
			Data:    nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}
	c.awardService.DeleteAward(id)
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Delete Award",
		Data:    nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}