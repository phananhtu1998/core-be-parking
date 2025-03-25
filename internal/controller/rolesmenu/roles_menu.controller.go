package rolesmenu

import (
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

var RolesMenus = new(cRolesMenu)

type cRolesMenu struct {
	rolesMenuService service.IRolesMenu
}

// CreateRolesMenu
// @Summary Create a new role menu mapping
// @Description Create a new mapping between roles and menus in the system
// @Tags RolesMenu
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param payload body model.RolesMenu true "Role menu mapping details"
// @Success 200 {object} response.ResponseData
// @Failure 500 {object} response.ErrorResponseData "Server error"
// @Router /rolesmenu/create_roles_menu [post]
func (c *cRolesMenu) CreateRolesMenu(ctx *gin.Context) {
	// Lấy dữ liệu từ request body
	var input model.RolesMenu
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeRoleMenuError, err.Error())
		return
	}

	// Gọi service để tạo roles menu
	code, result, err := service.RolesMenuItem().CreateRolesMenu(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	// Trả về kết quả thành công
	response.SuccessResponse(ctx, code, result)
}

// GetRoleMenuByRoleId
// @Summary      Lấy role menu theo ID
// @Description  API này trả về role menu theo ID
// @Tags         RolesMenu
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path   string  true  "ID role"
// @Param        search query string false "Từ khóa tìm kiếm"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /rolesmenu/get_role_menu_by_role_id/{id} [GET]
func (c *cRolesMenu) GetRoleMenuByRoleId(ctx *gin.Context) {
	id := ctx.Param("id")
	search := ctx.Query("search")
	code, menu, err := service.RolesMenuItem().GetRoleMenuByRoleId(ctx, id, search)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, menu)
}

// UpdateRolesMenu godoc
// @Summary      Cập nhật role menu
// @Description  Api cập nhật phân quyền menu cho role
// @Tags         RolesMenu
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "ID của role menu"
// @Param        payload body model.RolesMenu true "Thông tin cần cập nhật"
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /rolesmenu/update_roles_menu/{id} [put]
func (c *cRolesMenu) UpdateRolesMenu(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response.ErrorResponse(ctx, response.ErrCodeRoleNotFound, "ID không được để trống")
		return
	}

	var input model.RolesMenu
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeRoleNotFound, err.Error())
		return
	}

	code, result, err := service.RolesMenuItem().UpdateRolesMenu(ctx, id, &input)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, result)
}

// DeleteRolesMenu godoc
// @Summary      Xóa role menu
// @Description  Api xóa role menu
// @Tags         RolesMenu
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "ID của role menu"
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /rolesmenu/delete_roles_menu/{id} [delete]
func (c *cRolesMenu) DeleteRolesMenu(ctx *gin.Context) {
	id := ctx.Param("id")
	code, err := service.RolesMenuItem().DeleteRolesMenu(ctx, id)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}
	response.SuccessResponse(ctx, code, nil)
}
