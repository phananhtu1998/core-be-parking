package roleaccount

import (
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var RoleAccounts = new(cRoleaccount)

type cRoleaccount struct {
	roleAccountService service.IRoleAccount
}

// CreateRoleAccount
// @Summary Tạo role account
// @Description Api tạo role account cho hệ thống
// @Tags RoleAccount
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param payload body model.RoleAccount true "Role account mapping details"
// @Success 200 {object} response.ResponseData
// @Failure 500 {object} response.ErrorResponseData "Server error"
// @Router /roleaccount/create_roles_account [post]
func (c *cRoleaccount) CreateRoleAccount(ctx *gin.Context) {
	var params model.RoleAccount
	// check valid
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(response.ErrCodeParamInvalid, gin.H{"error": "Invalid input data"})
		return
	}
	code, result, err := service.RoleAccountItem().CreateRoleAccount(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}
	response.SuccessResponse(ctx, code, result)
}

// GetRoleAccountByRoleId
// @Summary      Lấy role account theo Role_Id
// @Description  API này trả về role account theo role_Id
// @Tags         RoleAccount
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path   string  true  "ID role"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /roleaccount/get_role_account_by_role_id/{id} [GET]
func (c *cRoleaccount) GetAllRoleAccountByRoleId(ctx *gin.Context) {
	Id := ctx.Param("id")
	// check uuid
	if _, err := uuid.Parse(Id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role id"})
		return
	}
	code, result, err := service.RoleAccountItem().GetAllRoleAccountByRoleId(ctx, Id)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}
	response.SuccessResponse(ctx, code, result)
}
