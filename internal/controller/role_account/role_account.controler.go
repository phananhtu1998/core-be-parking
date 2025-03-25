package roleaccount

import (
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

var RoleAccounts = new(cRoleaccount)

type cRoleaccount struct {
	roleAccountService service.IRoleAccount
}

// CreateRoleAccount
// @Summary Create a new role account mapping
// @Description Create a new mapping between roles and account in the system
// @Tags RoleAccount
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param payload body model.RoleAccount true "Role account mapping details"
// @Success 200 {object} response.ResponseData
// @Failure 500 {object} response.ErrorResponseData "Server error"
// @Router /roleaccount/create_roles_account [post]
func (c *cRoleaccount) CreateRoleAccount(ctx *gin.Context) {
	var input model.RoleAccount
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeRoleMenuError, err.Error())
		return
	}

	// Gọi service để tạo roles menu
	code, result, err := service.RoleAccountItem().CreateRoleAccount(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}
	// Trả về kết quả thành công
	response.SuccessResponse(ctx, code, result)
}
