package role

import (
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

var Roles = new(cRole)

type cRole struct {
	roleService service.IRole
}

// CreateRole godoc
// @Summary Create a new role
// @Description Create a new role with binary tree structure
// @Tags Role
// @Accept json
// @Produce json
// @Param role body model.RoleInput true "Role information"
// @Success 200 {object} model.RoleOutput
// @Router /role/create_role [post]
func (c *cRole) CreateRole(ctx *gin.Context) {
	var params model.RoleInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	codeRs, dataRs, err := c.roleService.CreateRole(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeRs, dataRs)
}
