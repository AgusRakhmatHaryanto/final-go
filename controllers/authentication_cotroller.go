package controllers

import (
	"final-project/data/request"
	"final-project/data/response"
	"final-project/helper"
	"final-project/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	authenticationService services.AuthenticationService
}

func NewAuthenticationController(service services.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{
		authenticationService: service,
	}
}

func (c *AuthenticationController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBind(&loginRequest)
	helper.ErrorPanic(err)

	token, err_token := c.authenticationService.Login(loginRequest)

	if err_token != nil {
		weResponse := response.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD REQUEST",
			Message: fmt.Sprintf("%v", err_token),
		}

		ctx.JSON(http.StatusBadRequest, weResponse)
		return
	}
	res := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Login Success",
		Data:    res,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *AuthenticationController) Register(ctx *gin.Context) {
	registerRequest := request.CreateNewUserRequest{}
	err := ctx.ShouldBind(&registerRequest)
	helper.ErrorPanic(err)
	c.authenticationService.Register(registerRequest)

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Register Success",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
