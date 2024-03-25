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
	var loginRequest request.LoginRequest
	err := ctx.ShouldBind(&loginRequest)
	helper.ErrorPanic(err)

	token, err_token := c.authenticationService.Login(loginRequest)

	if err_token != nil {
		weResponse := response.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD REQUEST",
			Message: fmt.Sprintf("%v", err_token),
		}

		ctx.Header("Content-Type", "application/json")
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

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *AuthenticationController) Register(ctx *gin.Context) {
	var registerRequest request.RegisterNewUserRequest
	err := ctx.ShouldBind(&registerRequest)
	helper.ErrorPanic(err)
	// if registerRequest.Username == "" || registerRequest.Email == "" || registerRequest.Password == "" {
	// 	webResponse := response.WebResponse{
	// 		Code:    http.StatusBadRequest,
	// 		Status:  "BAD REQUEST",
	// 		Message: "Username, Email, and Password cannot be empty",
	// 		Data:    nil,
	// 	}
	// 	ctx.Header("Content-Type", "application/json")
	// 	ctx.JSON(http.StatusBadRequest, webResponse)
	// 	return
	// }
	// email_req := c.authenticationService.FindByEmail(registerRequest.Email)
	// if email_req.ID != 0 {
	// 	webResponse := response.WebResponse{
	// 		Code:    http.StatusBadRequest,
	// 		Status:  "BAD REQUEST",
	// 		Message: "Email Already Registered",
	// 		Data:    nil,
	// 	}
	// 	ctx.Header("Content-Type", "application/json")
	// 	ctx.JSON(http.StatusBadRequest, webResponse)
	// 	return
	// }

	// username_req := c.authenticationService.FindByUsername(registerRequest.Username)
	// if username_req.ID != 0 {
	// 	webResponse := response.WebResponse{
	// 		Code:    http.StatusBadRequest,
	// 		Status:  "BAD REQUEST",
	// 		Message: "Username Already Registered",
	// 		Data:    nil,
	// 	}
	// 	ctx.Header("Content-Type", "application/json")
	// 	ctx.JSON(http.StatusBadRequest, webResponse)
	// 	return
	// }

	
	c.authenticationService.Register(request.RegisterNewUserRequest{
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
		Role:     "user",
	})

	// user := c.authenticationService
	// res_regis := response.NewRegisterResponse()

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Register Success",
		Data:    nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
