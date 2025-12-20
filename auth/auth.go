package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest/utils"
)

func Authentificate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token found"})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
