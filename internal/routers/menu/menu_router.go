package menu

import (
	"go-backend-api/internal/controller/menu"

	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

func (ar *MenuRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// private router
	menuRouterPrivate := Router.Group("/menu")
	//menuRouterPrivate.Use(middlewares.AuthenMiddleware())
	{
		menuRouterPrivate.POST("/create_menu", menu.Menus.CreateMenu)
		menuRouterPrivate.GET("/get_all_menu", menu.Menus.GetAllMenu)
		menuRouterPrivate.GET("/get_menu_by_id/:id", menu.Menus.GetMenuById)
		menuRouterPrivate.PUT("/update_multiple_menu/", menu.Menus.EditMenuById)
		menuRouterPrivate.DELETE("/delete/:id", menu.Menus.DeleteMenu)
	}

}
