package middleware

import (
	"final-project/config"
	"final-project/helper"
	"final-project/repository"
	"final-project/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(userRepository repository.UsersRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: Implement middleware
		var token string
		authHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"status":  "failed",
				"message": "Unauthorized",
			})
			return
		}

		config, _ := config.LoadConfig(".")
		sub, err := utils.ValidateToken(token, config.TokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"status":  "failed",
				"message": err.Error(),
			})
			return
		}

		id, err_id := strconv.Atoi(fmt.Sprint(sub))
		helper.ErrorPanic(err_id)
		result, err_result := userRepository.FindById(id)
		if err_result != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"status":  "failed",
				"message": err_result.Error(),
			})
			return
		}

		ctx.Set("currentUser", result)
		ctx.Next()
	}
}

func IsRole(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: Implement middleware
		id, role, err_extract := utils.ExtractToken(ctx)
		if err_extract != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"status":  "failed",
				"message": err_extract.Error(),
			})
			return
		}

		if role != role {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"status":  "failed",
				"message": "Unauthorized",
			})
			return
		}

		ctx.Set("currentUser", id)
		ctx.Next()
	}
}
