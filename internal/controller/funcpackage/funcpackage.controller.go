package funcpackage

import (
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

var Funcpackages = new(cFuncpackage)

type cFuncpackage struct {
	funcpackageService service.Ifuncpackage
}

// Funcpackages
// @Summary Tạo gói chức năng
// @Description Api tạo gói chức năng trong hệ thống
// @Tags Func packages
// @Accept json
// @Produce json
// @Security     BearerAuth
// @Security     ApiKeyAuth
// @Param        payload body model.FuncpackageInput true "payload"
// @Success 200 {object} response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router /funcpackage/create_func_package [post]
func (c *cFuncpackage) CreateFuncPackage(ctx *gin.Context) {
	var params model.FuncpackageInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	codeRole, dataRole, err := service.FuncpackageItem().CreateFuncPackage(ctx.Request.Context(), &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeRole, dataRole)
}

// Funcpackages
// @Summary Lấy gói chức năng
// @Description Api lấy gói chức năng trong hệ thống
// @Tags Func packages
// @Accept json
// @Produce json
// @Security     BearerAuth
// @Security     ApiKeyAuth
// @Success 200 {object} response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router /funcpackage/get_all_func_package [GET]
func (c *cFuncpackage) GetAllFuncPackage(ctx *gin.Context) {
	codeRole, dataRole, err := service.FuncpackageItem().GetAlFuncPackageByCreatedBy(ctx.Request.Context())
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeRole, dataRole)
}
