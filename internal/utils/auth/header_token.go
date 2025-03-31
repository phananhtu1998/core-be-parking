package auth

import (
	"log"

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
func ExtractRefreshToken(c *gin.Context) (string, bool) {
	refreshToken := c.GetHeader("RefreshToken")
	if refreshToken == "" {
		return "", false
	}
	return refreshToken, true
}

func ExtracapiKeyToken(c *gin.Context) (string, bool) {
	licenseHeader := c.GetHeader("X-API-Key")
	if licenseHeader == "" {
		return "", false
	}
	log.Println("licenseHeader: ", licenseHeader)
	// Kiểm tra nếu token có dạng "Bearer {token}"
	// if strings.HasPrefix(authHeader, "Bearer ") {
	// 	return strings.TrimPrefix(authHeader, "Bearer "), true
	// }
	// Trả về token thô nếu không có "Bearer "
	return licenseHeader, true
}
