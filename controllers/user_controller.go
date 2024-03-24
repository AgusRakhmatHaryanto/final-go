package controllers

import (
	"final-project/data/request"
	"final-project/data/response"
	"final-project/helper"
	"final-project/services"
	// "final-project/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UsersController struct {
	userService services.UsersService
}

func NewUsersController(service services.UsersService) *UsersController {
	return &UsersController{
		userService: service,
	}
}
func (c *UsersController) FindAllUsers(ctx *gin.Context) {
	log.Info().Msg("Get All Users")
	users := c.userService.GetAllUsers()
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Get All Users",
		Data:    users,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (c *UsersController) FindUserById(ctx *gin.Context) {
	log.Info().Msg("Get User By Id")
	userId := ctx.Param("id")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	user_res := c.userService.FindUserById(id)
	if user_res.ID == 0 {
		webResponse := response.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT FOUND",
			Message: "User Not Found",
			Data:    nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Get User By Id",
		Data:    user_res,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *UsersController) GetUserByEmail(ctx *gin.Context) {
	log.Info().Msg("Get User By Email")
	email := ctx.Param("email")
	user_res := c.userService.FindUserByEmail(email)
	if user_res.ID == 0 {
		webResponse := response.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT FOUND",
			Message: "User Not Found",
			Data:    nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Get User By Email",
		Data:    user_res,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *UsersController) GetUserByUsername(ctx *gin.Context) {
	log.Info().Msg("Get User By Username")
	username := ctx.Param("username")
	user_res := c.userService.FindUserByUsername(username)
	if user_res.ID == 0 {
		webResponse := response.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT FOUND",
			Message: "User Not Found",
			Data:    nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Get User By Username",
		Data:    user_res,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *UsersController) UpdateUser(ctx *gin.Context) {
	log.Info().Msg("Update User")
	userRequest := request.UpdateUserRequest{}
	err := ctx.ShouldBind(&userRequest)
	helper.ErrorPanic(err)

	userId := ctx.Param("id")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	userRequest.ID = id

	user := c.userService.FindUserById(id)
	if user.ID == 0 {
		webResponse := response.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT FOUND",
			Message: "User Not Found",
			Data:    nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}
	c.userService.UpdateUser(userRequest)

	data_res := response.UpdateUserResponse{
		ID:       id,
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Updated User",
		Data:    data_res,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *UsersController) DeleteUser(ctx *gin.Context) {
	log.Info().Msg("Delete User")
	userId := ctx.Param("id")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	c.userService.DeleteUser(id)
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Deleted User",
		Data:    nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *UsersController) CreateUser(ctx *gin.Context) {
	log.Info().Msg("Create User")
	userRequest := request.CreateNewUserRequest{}
	err := ctx.ShouldBind(&userRequest)
	helper.ErrorPanic(err)
	// if userRequest.Email == "" || userRequest.Password == "" || userRequest.Username == ""  {
	// 	webResponse := response.WebResponse{
	// 		Code:    http.StatusBadRequest,
	// 		Status:  "BAD REQUEST",
	// 		Message: "Email, Password, and Username cannot be empty",
	// 		Data:    nil,
	// 	}
	// 	ctx.Header("Content-Type", "application/json")
	// 	ctx.JSON(http.StatusBadRequest, webResponse)
	// 	return
	// }

	// username_req := c.userService.FindUserByUsername(userRequest.Username)
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

	// email_req := c.userService.FindUserByEmail(userRequest.Email)
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

	c.userService.CreateUser(userRequest)
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully Created User",
		Data:    nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
