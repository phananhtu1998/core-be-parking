package role

// import (
// 	"go-backend-api/internal/controller/role"
// 	"go-backend-api/internal/middlewares"

// 	"github.com/gin-gonic/gin"
// )

// type RoleRouter struct{}

// func (ar *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
// 	roleRouterPrivate := Router.Group("/role")
// 	roleRouterPrivate.Use(middlewares.AuthenMiddleware())
// 	roleRouterPrivate.Use(middlewares.RateLimiterPrivateMiddlewareRedis())
// 	roleRouterPrivate.Use(middlewares.LicenseMiddleware())
// 	//roleRouterPrivate.Use(middlewares.PermissionMiddleware(global.Enforcer))
// 	{
// 		roleRouterPrivate.POST("/create_role", role.Roles.CreateRole)
// 		roleRouterPrivate.GET("/get_all_roles", role.Roles.GetAllRoles)
// 		roleRouterPrivate.GET("/get_role_by_id/:id", role.Roles.GetRoleById)
// 		roleRouterPrivate.DELETE("/delete_role/:id", role.Roles.DeleteRole)
// 	}
// }
