package role

import (
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"
	"log"

	"github.com/gin-gonic/gin"
)

var Roles = new(cRole)

type cRole struct {
	roleService service.IRole
}

// Role
// @Summary Tạo role
// @Description Api tạo role trong hệ thống
// @Tags Role
// @Accept json
// @Produce json
// @Security     BearerAuth
// @Param        payload body model.Role true "payload"
// @Success 200 {object} response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router /role/create_role [post]
func (c *cRole) CreateRole(ctx *gin.Context) {
	var params model.Role
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	codeRole, dataRole, err := service.RoleItem().CreateRole(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeRole, dataRole)
}

// GetAllRoles
// @Summary Lấy danh sách role
// @Description Api lấy danh sách role trong hệ thống
// @Tags Role
// @Accept json
// @Produce json
// @Security     BearerAuth
// @Success 200 {object} response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router /role/get_all_roles [get]
func (c *cRole) GetAllRoles(ctx *gin.Context) {
	codeRole, dataRole, err := service.RoleItem().GetAllRoles(ctx)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeRole, dataRole)
}

// GetRoleById
// @Summary      Lấy role theo ID
// @Description  API này trả về role theo ID
// @Tags         Role
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path   string  true  "ID role cần lấy"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /role/get_role_by_id/{id} [GET]
func (c *cRole) GetRoleById(ctx *gin.Context) {
	id := ctx.Param("id")
	code, menu, err := service.RoleItem().GetRoleById(ctx, id)
	if err != nil {
		log.Printf("Error getting menu: %v", err)
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, menu)
}
