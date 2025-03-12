package account

import (
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"
	"log"

	"github.com/gin-gonic/gin"
)

var Accounts = new(cAccount)

type cAccount struct {
	accountService service.Iaccount
}

// GetAllAccount
// @Summary      Lấy danh sách tất cả tài khoản
// @Description  API này trả về danh sách tất cả tài khoản trong hệ thống
// @Tags         account management
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /admin/get_all_account [GET]
func (ac *cAccount) GetAllAccount(ctx *gin.Context) {
	code, accounts, err := service.AccountItem().GetAllAccount(ctx)
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
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        id   path   string  true  "ID tài khoản cần lấy"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /admin/get_account_by_id/{id} [GET]
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
