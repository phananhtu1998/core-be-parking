package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtracBearerToken(c *gin.Context) (string, bool) {
	//   Authorization: bearer token
	authHeader := c.GetHeader("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer "), true
	}
	return "", false
}
