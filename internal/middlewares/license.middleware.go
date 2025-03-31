package middlewares

import (
	"go-backend-api/internal/utils/auth"
	"go-backend-api/pkg/response"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LicenseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get request url path
		url := c.Request.URL.Path
		log.Println("url request", url)
		x_api_key, ok := auth.ExtracapiKeyToken(c)
		if !ok {
			log.Println("API key is missing")
			c.AbortWithStatusJSON(401, gin.H{"code": response.ErrUnauthorized, "err": "Unauthorized x-api-key", "description": ""})
			return
		}
		claims, err := auth.ParseJwtTokenPayload(x_api_key)
		if err != nil {
			log.Println("Invalid token")
			c.AbortWithStatusJSON(401, gin.H{"code": response.ErrUnauthorized, "err": "invalid x-api-key", "description": ""})
			return
		}
		if err != nil {
			log.Println("Invalid token")
			c.AbortWithStatusJSON(401, gin.H{
				"code":    response.ErrUnauthorized,
				"message": "invalid x-api-key",
			})
			return
		}

		// Kiểm tra trường "dateend" có tồn tại không
		dateendRaw, exists := claims["dateend"]
		if !exists {
			log.Println("Missing 'dateend' field in token")
			c.AbortWithStatusJSON(401, gin.H{
				"code":    response.ErrUnauthorized,
				"message": "missing expiration date",
			})
			return
		}
		// Chuyển "dateend" về kiểu string
		dateendStr, ok := dateendRaw.(string)
		if !ok {
			log.Println("Invalid 'dateend' format")
			c.AbortWithStatusJSON(401, gin.H{
				"code":    400,
				"message": "invalid expiration date format",
			})
			return
		}
		dateend, err := time.Parse("2006-01-02 15:04:05", dateendStr)
		if err != nil {
			log.Println("Error parsing 'dateend':", err)
			c.AbortWithStatusJSON(401, gin.H{
				"code":    400,
				"message": "invalid expiration date format",
			})
			return
		}

		// So sánh với thời gian hiện tại
		if time.Now().After(dateend) {
			log.Println("License expired")
			c.AbortWithStatusJSON(401, gin.H{
				"code":    400,
				"message": "License expired",
			})
			return
		}
		c.Next()
	}
}
