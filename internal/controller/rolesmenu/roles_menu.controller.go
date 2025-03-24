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
