package roleaccount

import (
	"bytes"
	"encoding/json"
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"
	"io"
	"net/http"

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
	// Đọc toàn bộ body request
	var params model.RoleAccount
	// Đọc raw JSON từ body
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	// Decode JSON với DisallowUnknownFields để phát hiện field dư
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid field in request"})
		return
	}
	// Bind JSON vào struct
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid field in request"})
		return
	}
	code, result, err := service.RoleAccountItem().CreateRoleAccount(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}
	response.SuccessResponse(ctx, code, result)
}
