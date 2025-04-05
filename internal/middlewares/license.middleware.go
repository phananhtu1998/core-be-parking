package middlewares

import (
	"context"
	"go-backend-api/internal/utils/auth"
	"go-backend-api/pkg/response"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LicenseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get request URL path
		url := c.Request.URL.Path
		log.Println("URL request:", url)

		// Lấy x-api-key từ header
		x_api_key, ok := auth.ExtracapiKeyToken(c)
		if !ok {
			log.Println("API key is missing")
			c.AbortWithStatusJSON(401, gin.H{
				"code":    response.ErrUnauthorized,
				"message": "Unauthorized x-api-key",
			})
			return
		}

		// Parse JWT token
		claims, err := auth.ParseJwtTokenPayload(x_api_key)
		if err != nil {
			log.Println("Invalid token")
			c.AbortWithStatusJSON(401, gin.H{
				"code":    response.ErrUnauthorized,
				"message": "Invalid x-api-key",
			})
			return
		}

		// Kiểm tra trường "dateend"
		dateendRaw, exists := claims["dateend"]
		if !exists {
			log.Println("Missing 'dateend' field in token")
			c.AbortWithStatusJSON(401, gin.H{
				"code":    response.ErrUnauthorized,
				"message": "Key đã hết hạn vui lòng gia hạn lại key!",
			})
			return
		}

		// Chuyển "dateend" về kiểu string
		dateendStr, ok := dateendRaw.(string)
		if !ok {
			log.Println("Invalid 'dateend' format")
			c.AbortWithStatusJSON(400, gin.H{
				"code":    400,
				"message": "Invalid expiration date format",
			})
			return
		}

		// Store dateend as string in context
		ctx := context.WithValue(c.Request.Context(), "dateend", dateendStr)
		c.Request = c.Request.WithContext(ctx)

		// Nếu giá trị là "NO_EXPIRATION", bỏ qua kiểm tra hết hạn
		if dateendStr == "NO_EXPIRATION" {
			log.Println("License is set to NO_EXPIRATION, skipping expiration check")
			c.Next()
			return
		}

		// Chuyển đổi dateend thành kiểu time.Time
		dateend, err := time.Parse("2006-01-02 15:04:05", dateendStr)
		if err != nil {
			log.Println("Error parsing 'dateend':", err)
			c.AbortWithStatusJSON(400, gin.H{
				"code":    400,
				"message": "Invalid expiration date format",
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
		// Nếu hợp lệ, tiếp tục request
		c.Next()
	}
}
