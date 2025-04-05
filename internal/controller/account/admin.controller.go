package account

import (
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Accounts = new(cAccount)

type cAccount struct {
	accountService service.Iaccount
}

// GetAllAccount
// @Summary      Lấy danh sách tất cả tài khoản
// @Description  API này trả về danh sách tất cả tài khoản trong hệ thống
// @Tags         Account
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Security     ApiKeyAuth
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /account/get_all_account [GET]
func (ac *cAccount) GetAllAccount(ctx *gin.Context) {
	code, accounts, err := service.AccountItem().GetAllAccount(ctx.Request.Context())
	if err != nil {
		log.Printf("Error getting account: %v", err)
		response.ErrorResponse(ctx, code, err.Error())
		return
	}
	response.SuccessResponse(ctx, code, accounts)
}

// GetAccountById
// @Summary      Lấy tài khoản theo ID
// @Description  API này trả về tài khoản theo ID
// @Tags         Account
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Security     ApiKeyAuth
// @Param        id   path   string  true  "ID tài khoản cần lấy"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /account/get_account_by_id/{id} [GET]
func (ac *cAccount) GetAccountById(ctx *gin.Context) {
	id := ctx.Param("id") // Lấy ID từ request
	code, account, err := service.AccountItem().GetAccountById(ctx, id)
	if err != nil {
		log.Printf("Error getting account: %v", err)
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, account)
}

// UpdateAccount
// @Summary      Cập nhật tài khoản
// @Description  API này cập nhật thông tin tài khoản dựa trên ID
// @Tags         Account
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Security     ApiKeyAuth
// @Param        id   path   string  true  "ID tài khoản cần cập nhật"
// @Param        body body   model.AccountInput true "Dữ liệu cập nhật tài khoản"
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /account/update_account/{id} [PUT]
func (ac *cAccount) UpdateAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	var modelAccount model.AccountInput
	if err := ctx.ShouldBindJSON(&modelAccount); err != nil {
		ctx.JSON(response.ErrCodeParamInvalid, gin.H{"error": "Invalid input data"})
		return
	}
	code, account, err := service.AccountItem().UpdateAccount(ctx, &modelAccount, id)
	if err != nil {
		log.Printf("Error getting account: %v", err)
		response.ErrorResponse(ctx, code, err.Error())
		return
	}
	response.SuccessResponse(ctx, code, account)
}

// DeleteAccount
// @Summary      Xóa tài khoản
// @Description  API này xóa tài khoản dựa trên ID
// @Tags         Account
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Security     ApiKeyAuth
// @Param        id   path   string  true  "ID của tài khoản cần xóa"
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /account/delete_account/{id} [DELETE]
func (ac *cAccount) DeleteAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	codeResult, err := service.AccountItem().DeleteAccount(ctx, id)
	if err != nil {
		log.Printf("Error getting account: %v", err)
		response.ErrorResponse(ctx, codeResult, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeResult, err)
}

// CreateAccount
// @Summary      Tạo tài khoản mới
// @Description  API này cho phép tạo tài khoản mới
// @Tags         Account
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Security     ApiKeyAuth
// @Param        body  body   model.AccountInput  true  "Thông tin tài khoản cần tạo"
// @Success      200   {object}  response.ResponseData
// @Failure      400   {object}  response.ErrorResponseData
// @Failure      500   {object}  response.ErrorResponseData
// @Router       /account/create_account [POST]
func (ac *cAccount) CreateAccount(ctx *gin.Context) {
	var params model.AccountInput
	//  Lấy role account của tài khoản đang tạo
	// Kiểm tra số lượng account được phép tạo
	// check valid
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// call service CreateAccount
	code, account, err := service.AccountItem().CreateAccount(ctx.Request.Context(), &params)
	if err != nil {
		log.Printf("Error creating account: %v", err)
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	// respone data
	response.SuccessResponse(ctx, code, account)
}
