package middlewares

import (
	"context"
	"go-backend-api/internal/utils/auth"
	"go-backend-api/pkg/response"
	"log"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get request url path
		url := c.Request.URL.Path
		log.Println("url request", url)
		// check header authentication
		jwtToken, valid := auth.ExtracBearerToken(c)
		if !valid {
			log.Println("Token is missing")
			c.AbortWithStatusJSON(401, gin.H{"code": response.ErrUnauthorized, "err": "Unauthorized", "description": ""})
			return
		}
		// validate jwt token
		claims, err := auth.VerifyToken(jwtToken)
		if err != nil {
			log.Println("Invalid token")
			c.AbortWithStatusJSON(401, gin.H{"code": response.ErrUnauthorized, "err": "invalid token", "description": ""})
			return
		}
		// update claims to context
		log.Println("claims::: UUID::", claims.Subject)
		ctx := context.WithValue(c.Request.Context(), "subjectUUID", claims.Subject)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// // AuthMiddleware kiểm tra Bearer Token và X-API-Key
// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Lấy Bearer Token từ Header
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Bearer token"})
// 			c.Abort()
// 			return
// 		}

// 		// Lấy token sau "Bearer "
// 		token := strings.TrimPrefix(authHeader, "Bearer ")
// 		if !validateToken(token) {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Bearer token"})
// 			c.Abort()
// 			return
// 		}

// 		// Lấy X-API-Key từ Header
// 		apiKey := c.GetHeader("X-API-Key")
// 		if apiKey == "" || !validateAPIKey(apiKey) {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API Key"})
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }

// // validateToken kiểm tra tính hợp lệ của Bearer Token
// func validateToken(token string) bool {
// 	// Thay "valid-token" bằng logic xác thực thực tế (VD: JWT)
// 	return token == "valid-token"
// }

// // validateAPIKey kiểm tra tính hợp lệ của X-API-Key
// func validateAPIKey(key string) bool {
// 	// Thay "valid-api-key" bằng API Key thực tế từ DB hoặc file cấu hình
// 	return key == "valid-api-key"
// }
