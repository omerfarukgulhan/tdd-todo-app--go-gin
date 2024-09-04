package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"todo-app--go-gin/common/util/security"
)

func Authenticate(context *gin.Context) {
	authHeader := context.Request.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	userId, userEmail, err := security.ValidateToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	context.Set("userId", userId)
	context.Set("userEmail", userEmail)
	context.Next()
}
