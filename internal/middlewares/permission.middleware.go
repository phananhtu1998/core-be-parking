package middlewares

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// PermissionMiddleware kiểm tra quyền truy cập dựa trên RBAC
func PermissionMiddleware(enforcer *casbin.SyncedEnforcer, path string, method string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Kiểm tra quyền truy cập với role mặc định (ví dụ: "anonymous")
		allowed, err := enforcer.Enforce("anonymous", path, method)
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
