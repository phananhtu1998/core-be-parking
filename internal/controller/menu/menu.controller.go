package menu

import (
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

var Menus = new(cMenu)

type cMenu struct {
	loginService service.IMenu
}

// Menu
// @Summary      Tạo menu
// @Description  API tạo menu trong hệ thống
// @Tags         Menu
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Security     ApiKeyAuth
// @Param        payload body model.MenuInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /menu/create_menu [post]
func (c *cMenu) CreateMenu(ctx *gin.Context) {
	// Implement logic for create
	var params model.MenuInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	codeMenu, dataMenu, err := service.MenuItem().CreateMenu(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeMenu, dataMenu)
}

// Menu
// @Summary      Lấy danh sách tất cả menu
// @Description  API này trả về danh sách tất cả menu trong hệ thống
// @Tags         Menu
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Security     ApiKeyAuth
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /menu/get_all_menu [GET]
func (ac *cMenu) GetAllMenu(ctx *gin.Context) {
	code, menus, err := service.MenuItem().GetAllMenu(ctx)
	if err != nil {
		log.Printf("Error getting menu: %v", err)
		response.ErrorResponse(ctx, code, err.Error())
		return
	}
	response.SuccessResponse(ctx, code, menus)
}

// Menu
// @Summary      Lấy menu theo ID
// @Description  API này trả về menu theo ID
// @Tags         Menu
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Security     ApiKeyAuth
// @Param        id   path   string  true  "ID menu cần lấy"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /menu/get_menu_by_id/{id} [GET]
func (ac *cMenu) GetMenuById(ctx *gin.Context) {
	id := ctx.Param("id")
	code, menu, err := service.MenuItem().GetMenuById(ctx, id)
	if err != nil {
		log.Printf("Error getting menu: %v", err)
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, menu)
}

// UpdateMenu
// @Summary      Cập nhật nhiều menu
// @Description  API này cập nhật danh sách menu dựa trên danh sách ID
// @Tags         Menu
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Security     ApiKeyAuth
// @Param        body body   []model.MenuInput true "Danh sách menu cần cập nhật"
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /menu/update_multiple_menu [PUT]
func (ac *cMenu) EditMenuById(ctx *gin.Context) {
	var menuUpdates []model.MenuInput
	if err := ctx.ShouldBindJSON(&menuUpdates); err != nil {
		ctx.JSON(response.ErrCodeParamInvalid, gin.H{"error": "Invalid input data"})
		return
	}

	// Lấy context chuẩn
	requestCtx := ctx.Request.Context()

	// Gọi service cập nhật nhiều menu
	code, updatedMenus, err := service.MenuItem().EditMenuById(requestCtx, menuUpdates)
	if err != nil {
		log.Printf("Error updating menu: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  code,
			"error": err.Error(),
		})
		return
	}

	response.SuccessResponse(ctx, code, updatedMenus)
}

// DeleteMenu
// @Summary      Xóa menu
// @Description  API này xóa một menu dựa trên ID
// @Tags         Menu
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Security     ApiKeyAuth
// @Param        id   path      string  true  "Menu ID"
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /menu/delete/{id} [DELETE]
func (ac *cMenu) DeleteMenu(ctx *gin.Context) {
	id := ctx.Param("id")
	// Gọi service xóa menu
	code, err := service.MenuItem().DeleteMenu(ctx.Request.Context(), id)
	if err != nil {
		log.Printf("Error deleting menu: %v", err)
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, nil)
}

// CreateMultipleMenus
// @Summary      Tạo nhiều menu
// @Description  API tạo nhiều menu cùng lúc trong hệ thống
// @Tags         Menu
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Security     ApiKeyAuth
// @Param        payload body []model.MenuInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /menu/create_multiple_menus [post]
func (c *cMenu) CreateMultipleMenus(ctx *gin.Context) {
	var params []model.MenuInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	codeMenu, dataMenu, err := service.MenuItem().CreateMultipleMenus(ctx, params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeMenu, dataMenu)
}
