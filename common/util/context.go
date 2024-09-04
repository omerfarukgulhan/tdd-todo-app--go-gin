package util

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func GetUserIdFromContext(ctx *gin.Context) (int, error) {
	userId, exists := ctx.Get("userId")
	if !exists {
		return 0, errors.New("User id not found in context")
	}

	userIdInt, ok := userId.(int)
	if !ok {
		return 0, errors.New("User id type assertion failed")
	}

	return userIdInt, nil
}
