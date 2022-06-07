package middlewares

import (
	"fmt"
	"hacktiv8-day12/helper"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "Bearer"
	authorizationKey        = "userData"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get(authorizationHeaderKey)
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "token not found",
			})
			return
		}

		fmt.Println(token)

		bearer := strings.HasPrefix(token, authorizationTypeBearer)
		if !bearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "no bearer",
			})
			return
		}

		tokenString := strings.Split(token, "Bearer ")[1]
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "token not found after bearer ",
			})
			return
		}

		claims, err := helper.VerifyToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": err.Error(),
			})
			return
		}
		fmt.Printf("%v", claims)

		var data = claims.(jwt.MapClaims)

		ctx.Set("email", data["email"])
		ctx.Set("id", data["id"])
		ctx.Next()
	}
}
