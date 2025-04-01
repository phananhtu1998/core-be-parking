package license

import (
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Licenses = new(cLicense)

type cLicense struct {
	LicenseService service.ILicense
}

// CreateLicense
// @Summary      Tạo mới license cho gói chức năng
// @Description  API tạo mới license cho gói chức năng
// @Tags         License
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Security     ApiKeyAuth
// @Param        body  body   model.License  true  "Thông tin License cần tạo"
// @Success      200   {object}  response.ResponseData
// @Failure      400   {object}  response.ErrorResponseData
// @Failure      500   {object}  response.ErrorResponseData
// @Router       /license/create_license [POST]
func (ac *cLicense) CreateAccount(ctx *gin.Context) {
	var params model.License
	// check valid
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	code, license, err := service.LicenseItem().CreateLicense(ctx.Request.Context(), &params)
	if err != nil {
		log.Printf("Error creating license: %v", err)
		response.ErrorResponse(ctx, code, err.Error())
		return
	}
	// respone data
	response.SuccessResponse(ctx, code, license)
}
