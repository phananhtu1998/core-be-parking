package rolesmenu

import (
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

var RolesMenu = new(cRolesMenu)

type cRolesMenu struct {
	rolesMenuService service.IRolesMenu
}

// Rolesmenu
// @Summary Tạo role
// @Description Api tạo role trong hệ thống
// @Tags Role
// @Accept json
// @Produce json
// @Security     BearerAuth
// @Param        payload body model.RolesMenu true "payload"
// @Success 200 {object} response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router /rolesmenu/create_roles_menu [post]

func (c *cRolesMenu) CreateRolesMenu(ctx *gin.Context) {
	// Lấy dữ liệu từ request body
	var input model.RolesMenu
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeRoleMenuError, err.Error())
		return
	}

	// Gọi service để tạo roles menu
	code, result, err := c.rolesMenuService.CreateRolesMenu(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	// Trả về kết quả thành công
	response.SuccessResponse(ctx, code, result)
}
