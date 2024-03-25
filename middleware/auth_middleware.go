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
		id_string := strconv.Itoa(result.ID)

		ctx.Set("currentUser", gin.H{
			"id":    id_string,
			"role":  result.Role,
		})
		ctx.Next()
	}
}

func IsRole(roleType string) gin.HandlerFunc {
    return func(c *gin.Context) {
        // TODO: Implement middleware
        id, role, err_extract := utils.ExtractToken(c)
        if err_extract != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "code":    http.StatusUnauthorized,
                "status":  "failed",
                "message": err_extract.Error(),
            })
            return
        }

        if role != roleType { // Menggunakan roleType dari parameter fungsi
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "code":    http.StatusUnauthorized,
                "status":  "failed",
                "message": "Unauthorized",
            })
            return
        }
        c.Set("currentUser", id)
        c.Next()
    }
}


