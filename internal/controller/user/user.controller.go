package user

import (
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Users = new(cUser)

type cUser struct {
	userService service.Iuser
}

// tạo người dùng thì truyền role_id
// CreateUser
// @Summary      Tạo tài người dùng
// @Description  API này cho phép tạo người dùng mới
// @Tags         User
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Security     ApiKeyAuth
// @Param        body  body   model.AccountInput  true  "Thông tin người dùng cần tạo"
// @Success      200   {object}  response.ResponseData
// @Failure      400   {object}  response.ErrorResponseData
// @Failure      500   {object}  response.ErrorResponseData
// @Router       /user/create_user [POST]
func (ac *cUser) CreateUser(ctx *gin.Context) {
	var params model.AccountInput
	// check valid
	if err := ctx.ShouldBindJSON(&params); err != nil {
		log.Printf("Binding error: %+v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	// call service CreateUser
	code, account, err := service.UserItem().CreateUser(ctx.Request.Context(), &params)
	if err != nil {
		log.Printf("Error creating account: %v", err)
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	// respone data
	response.SuccessResponse(ctx, code, account)
}
