package menu

import (
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

var Menus = new(cMenu)

type cMenu struct {
	loginService service.IMenu
}

// Menu
// @Summary      Menu
// @Description  Menu
// @Tags         Menu
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        payload body model.MenuInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /menu/create_menu [post]
func (c *cMenu) CreateMenu(ctx *gin.Context) {
	// Implement logic for create
	var params model.MenuInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	codeRs, dataRs, err := service.MenuItem().CreateMenu(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeRs, dataRs)
}
