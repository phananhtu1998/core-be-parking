package login

import (
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"
	"log"

	"github.com/gin-gonic/gin"
)

var Logins = new(cLogin)

type cLogin struct {
	loginService service.ILogin
}

// Authenticate
// @Summary      Login
// @Description  Login
// @Tags         Authenticate
// @Accept       json
// @Produce      json
// @Param        payload body model.LoginInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /auth/login [post]
func (c *cLogin) Login(ctx *gin.Context) {
	// Implement logic for login
	var params model.LoginInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	codeRs, dataRs, err := service.LoginItem().Login(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeRs, dataRs)
}

// Authenticate
// @Summary      Logout
// @Description  Logout
// @Tags         Authenticate
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /auth/logout [post]
func (c *cLogin) Logout(ctx *gin.Context) {
	codeRs, err := service.LoginItem().Logout(ctx.Request.Context()) // truyền ctx.Request.Context để truyền giá trị subjectUUID
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeRs, nil)
}

// Authenticate
// @Summary      RefreshToken
// @Description  RefreshToken
// @Tags         Authenticate
// @Accept       json
// @Produce      json
// @Param        RefreshToken  header  string  true  "Refresh Token"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /auth/refresh-token [post]
func (c *cLogin) RefreshTokens(ctx *gin.Context) {
	codeRs, data, err := service.LoginItem().RefreshTokens(ctx.Request.Context()) // truyền ctx.Request.Context để truyền giá trị subjectUUID
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeRs, data)
}

// Authenticate
// @Summary      ChangePassword
// @Description  ChangePassword
// @Tags         Authenticate
// @Accept       json
// @Produce      json
// @Param        payload body model.ChangePasswordInput true "payload"
// @Security     BearerAuth
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /auth/change_password [post]
func (c *cLogin) ChangePassword(ctx *gin.Context) {
	var params model.ChangePasswordInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	log.Println("params:", &params)
	codeRs, dataRs, err := service.LoginItem().ChangePassword(ctx.Request.Context(), &params) // truyền ctx.Request.Context để truyền giá trị subjectUUID
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeRs, dataRs)
}
