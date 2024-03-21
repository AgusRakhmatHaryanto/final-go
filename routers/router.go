package routers

import (
	"final-project/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(authenticationController *controllers.AuthenticationController) *gin.Engine {
	service := gin.Default()
	service.GET("", func (ctx *gin.Context)  {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome Home"})
	})

	router  := service.Group("api/v1")

	authenticationRouter := router.Group("auth")
	{
		authenticationRouter.POST("login", authenticationController.Login)
		authenticationRouter.POST("register", authenticationController.Register)
	}
	return service
}
