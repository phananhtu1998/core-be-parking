package middlewares

import (
	"encoding/json"
	consts "go-backend-api/internal/const"
	"go-backend-api/internal/model"
	"go-backend-api/internal/utils/auth"
	"go-backend-api/internal/utils/cache"
	"go-backend-api/internal/utils/rbac"
	"net/http"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// Middleware kiểm tra quyền truy cập bằng Casbin
func PermissionMiddleware(enforcer *casbin.SyncedEnforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		// Lấy token từ request
		jwtToken, _ := auth.ExtracBearerToken(c)
		claims, _ := auth.VerifyTokenSubject(jwtToken)
		var infoUser model.GetCacheToken
		if err := cache.GetCache(ctx, claims.Subject, &infoUser); err != nil {
			return
		}

		// Lấy danh sách quyền của user từ DB
		lstUserPermission, err := rbac.GetFullPermisionByAccount(ctx, infoUser.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user permissions"})
			c.Abort()
			return
		}

		// Xóa quyền cũ để đảm bảo load lại đúng
		enforcer.DeletePermissionsForUser(infoUser.ID)

		// Load quyền vào Casbin
		for _, perm := range lstUserPermission {
			// Chuyển JSON string thành slice `[]string`
			perm.Method = strings.ReplaceAll(perm.Method, "'", "\"")
			var methods []string
			err := json.Unmarshal([]byte(perm.Method), &methods)
			if err != nil {
				continue
			}
			// Thêm từng method vào Casbin
			for _, method := range methods {
				enforcer.AddPermissionForUser(perm.Id, perm.Menu_group_name, method)
			}
		}

		// Kiểm tra quyền user với Casbin
		obj := strings.TrimPrefix(c.Request.URL.Path, consts.HOST_PREFIX)
		sub := infoUser.ID
		act := c.Request.Method

		allowed, err := enforcer.Enforce(sub, obj, act)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error checking permission"})
			return
		}

		if !allowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			return
		}

		c.Next()
	}
}
