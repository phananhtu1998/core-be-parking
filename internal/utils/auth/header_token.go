package auth

import (
	"github.com/gin-gonic/gin"
)

func ExtracBearerToken(c *gin.Context) (string, bool) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", false
	}
	// Kiểm tra nếu token có dạng "Bearer {token}"
	// if strings.HasPrefix(authHeader, "Bearer ") {
	// 	return strings.TrimPrefix(authHeader, "Bearer "), true
	// }
	// Trả về token thô nếu không có "Bearer "
	return authHeader, true
}
