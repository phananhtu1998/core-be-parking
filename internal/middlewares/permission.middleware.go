package middlewares

import (
	"encoding/json"
	consts "go-backend-api/internal/const"
	"go-backend-api/internal/model"
	"go-backend-api/internal/utils/auth"
	"go-backend-api/internal/utils/cache"
	"go-backend-api/internal/utils/rbac"
	"log"
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
			perm.Method = strings.ReplaceAll(perm.Method, "'", "\"")
			var methods []string
			err := json.Unmarshal([]byte(perm.Method), &methods)
			if err != nil {
				continue
			}
			for _, method := range methods {
				enforcer.AddPermissionForUser(perm.Id, perm.Menu_group_name, method)
			}
		}

		// Kiểm tra quyền user với Casbin
		fullPath := strings.TrimPrefix(c.Request.URL.Path, consts.HOST_PREFIX)
		pathParts := strings.Split(fullPath, "/")
		
		sub := infoUser.ID
		act := c.Request.Method
		allowed := false

		// Kiểm tra quyền trực tiếp với full path
		hasDirectPermission, err := enforcer.Enforce(sub, fullPath, act)
		if err == nil && hasDirectPermission {
			allowed = true
		}

		// Nếu không có quyền trực tiếp, kiểm tra quyền của các path cha
		if !allowed && len(pathParts) > 1 {
			// Xây dựng và kiểm tra từng path cha có thể
			var parentPath string
			for i := 0; i < len(pathParts)-1; i++ {
				if i == 0 {
					parentPath = pathParts[0]
				} else {
					parentPath = parentPath + "/" + pathParts[i]
				}
				
				hasParentPermission, err := enforcer.Enforce(sub, parentPath, act)
				if err == nil && hasParentPermission {
					allowed = true
					break
				}
			}
		}

		if !allowed {
			log.Printf("Permission denied for user %s, path %s, method %s", sub, fullPath, act)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			return
		}

		c.Next()
	}
}
