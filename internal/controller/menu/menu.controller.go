package menu

import (
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"
	"log"

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

	codeRs, dataRs, err := service.MenuItem().CreateMenu(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeRs, dataRs)
}

// Menu
// @Summary      Lấy danh sách tất cả menu
// @Description  API này trả về danh sách tất cả menu trong hệ thống
// @Tags         Menu
// @Accept       json
// @Produce      json
// @Security     BearerAuth
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
