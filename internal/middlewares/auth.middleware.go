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
		claims, err := auth.VerifyTokenSubject(jwtToken)
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
func AuthenMiddlewareV2() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get request url path
		url := c.Request.URL.Path
		log.Println("url request", url)
		// check header authentication
		refreshToken, valid := auth.ExtractRefreshToken(c)
		if !valid {
			log.Println("Token is missing")
			c.AbortWithStatusJSON(401, gin.H{"code": response.ErrUnauthorized, "err": "Unauthorized", "description": ""})
			return
		}
		// validate jwt token
		claims, err := auth.VerifyTokenSubject(refreshToken)
		if err != nil {
			log.Println("Invalid token")
			c.AbortWithStatusJSON(401, gin.H{"code": response.ErrUnauthorized, "err": "invalid token", "description": ""})
			return
		}
		// update claims to context
		//log.Println("claims::: UUID::", claims.Subject)
		newCtx := context.WithValue(c.Request.Context(), "subjectUUID", claims.Subject)
		newCtx = context.WithValue(newCtx, "refreshToken", refreshToken)
		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}
}
