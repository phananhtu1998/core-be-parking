package roleaccount

// type RoleAccountsRouterGroup struct {
// 	RoleAccountRouter
// }

// func (ar *RoleAccountRouter) InitRoleAccountsRouter(Router *gin.RouterGroup) {
// 	roleAccountRouterPrivate := Router.Group("/roleaccount")
// 	roleAccountRouterPrivate.Use(middlewares.AuthenMiddleware())
// 	roleAccountRouterPrivate.Use(middlewares.RateLimiterPrivateMiddlewareRedis())
// 	roleAccountRouterPrivate.Use(middlewares.LicenseMiddleware())
// 	{
// 		roleAccountRouterPrivate.POST("/create_roles_account", roleaccount.RoleAccounts.CreateRoleAccount)
// 		roleAccountRouterPrivate.GET("/get_role_account_by_role_id/:id", roleaccount.RoleAccounts.GetAllRoleAccountByRoleId)
// 		roleAccountRouterPrivate.GET("/get_role_account_by_account_id/:id", roleaccount.RoleAccounts.GetAllRoleAccountByAccountId)
// 		roleAccountRouterPrivate.PUT("/update_role_account/:id", roleaccount.RoleAccounts.UpdateRoleAccount)
// 		roleAccountRouterPrivate.DELETE("/delete_multiple_role_account", roleaccount.RoleAccounts.DeleteRoleAccount)
// 	}
// }
