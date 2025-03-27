package middlewares

import (
	"go-backend-api/internal/utils/auth"
	"go-backend-api/internal/utils/rbac"
	"go-backend-api/pkg/response"
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// PermissionMiddleware kiểm tra quyền truy cập bằng Casbin
func PermissionMiddleware(enforcer *casbin.SyncedEnforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		log.Println("Request URL:", c.Request.URL.Path)
		jwtToken, valid := auth.ExtracBearerToken(c)
		if !valid {
			log.Println("Token is missing")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": response.ErrUnauthorized, "err": "Unauthorized",
			})
			return
		}

		claims, err := auth.VerifyTokenSubject(jwtToken)
		if err != nil {
			log.Println("Invalid token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": response.ErrUnauthorized, "err": "Invalid token",
			})
			return
		}

		// Lấy danh sách quyền của user từ DB
		lstUserPermission, err := rbac.GetFullPermisionByAccount(ctx, claims.Subject)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user permissions"})
			c.Abort()
			return
		}

		// Load quyền vào Casbin nếu chưa có
		for _, perm := range lstUserPermission {
			enforcer.AddPermissionForUser(claims.Subject, perm.Menu_group_name, perm.Method)
		}

		sub := claims.Subject     // Người dùng
		obj := c.Request.URL.Path // Đối tượng truy cập (endpoint)
		act := c.Request.Method   // Hành động (GET, POST, DELETE,...)

		allowed, err := enforcer.Enforce(sub, obj, act)
		if err != nil {
			log.Println("Error checking permission:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error checking permission"})
			return
		}

		if !allowed {
			log.Println("Permission denied for:", sub, "on", obj, "with action", act)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			return
		}

		c.Next()
	}
}
