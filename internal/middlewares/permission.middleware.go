package middlewares

import (
	"go-backend-api/internal/utils/rbac"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// PermissionMiddleware kiểm tra quyền truy cập dựa trên RBAC
func PermissionMiddleware(enforcer *casbin.SyncedEnforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// path := c.Request.URL.Path
		// method := c.Request.Method
		// jwtToken, valid := auth.ExtracBearerToken(c)
		ctx := c.Request.Context()
		_, lstPermission, err := rbac.GetFullPermision(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving permissions"})
			c.Abort()
			return
		}
		allowed, err := enforcer.Enforce(lstPermission)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking permission"})
			c.Abort()
			return
		}

		if !allowed {
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}
