package roleaccount

// import (
// 	"fmt"
// 	"go-backend-api/internal/model"
// 	"go-backend-api/internal/service"
// 	"go-backend-api/pkg/response"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// var RoleAccounts = new(cRoleaccount)

// type cRoleaccount struct {
// 	roleAccountService service.IRoleAccount
// }

// // CreateRoleAccount
// // @Summary Tạo role account
// // @Description Api tạo role account cho hệ thống
// // @Tags RoleAccount
// // @Accept json
// // @Produce json
// // @Security BearerAuth
// // @Security     ApiKeyAuth
// // @Param payload body model.RoleAccount true "Role account mapping details"
// // @Success 200 {object} response.ResponseData
// // @Failure 500 {object} response.ErrorResponseData "Server error"
// // @Router /roleaccount/create_roles_account [post]
// func (c *cRoleaccount) CreateRoleAccount(ctx *gin.Context) {
// 	var params model.RoleAccount
// 	// check valid
// 	if err := ctx.ShouldBindJSON(&params); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
// 		return
// 	}
// 	code, result, err := service.RoleAccountItem().CreateRoleAccount(ctx, &params)
// 	if err != nil {
// 		response.ErrorResponse(ctx, code, err.Error())
// 		return
// 	}
// 	response.SuccessResponse(ctx, code, result)
// }

// // GetRoleAccountByRoleId
// // @Summary      Lấy role account theo Role_Id
// // @Description  API này trả về role account theo role_Id
// // @Tags         RoleAccount
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Security     ApiKeyAuth
// // @Param        id   path   string  true  "ID role"
// // @Success      200  {object}  response.ResponseData
// // @Failure      500  {object}  response.ErrorResponseData
// // @Router       /roleaccount/get_role_account_by_role_id/{id} [GET]
// func (c *cRoleaccount) GetAllRoleAccountByRoleId(ctx *gin.Context) {
// 	Id := ctx.Param("id")
// 	// check uuid
// 	if _, err := uuid.Parse(Id); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role id"})
// 		return
// 	}
// 	code, result, err := service.RoleAccountItem().GetAllRoleAccountByRoleId(ctx, Id)
// 	if err != nil {
// 		response.ErrorResponse(ctx, code, err.Error())
// 		return
// 	}
// 	response.SuccessResponse(ctx, code, result)
// }

// // GetAllRoleAccountByAccountId
// // @Summary      Lấy role account theo Account_Id
// // @Description  API này trả về role account theo Account_Id
// // @Tags         RoleAccount
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Security     ApiKeyAuth
// // @Param        id   path   string  true  "ID account"
// // @Success      200  {object}  response.ResponseData
// // @Failure      500  {object}  response.ErrorResponseData
// // @Router       /roleaccount/get_role_account_by_account_id/{id} [GET]
// func (c *cRoleaccount) GetAllRoleAccountByAccountId(ctx *gin.Context) {
// 	Id := ctx.Param("id")
// 	// check uuid
// 	if _, err := uuid.Parse(Id); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role id"})
// 		return
// 	}
// 	code, result, err := service.RoleAccountItem().GetAllRoleAccountByAccountId(ctx.Request.Context(), Id)
// 	if err != nil {
// 		response.ErrorResponse(ctx, code, err.Error())
// 		return
// 	}
// 	response.SuccessResponse(ctx, code, result)
// }

// // DeleteRoleAccountById
// // @Summary      Xóa nhiều Role Account
// // @Description  API này xóa danh sách Role Account dựa trên danh sách ID
// // @Tags         RoleAccount
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Security     ApiKeyAuth
// // @Param        ids body []string true "Danh sách ID cần xóa"
// // @Success      200  {object}  response.ResponseData
// // @Failure      400  {object}  response.ErrorResponseData
// // @Failure      500  {object}  response.ErrorResponseData
// // @Router       /roleaccount/delete_multiple_role_account [DELETE]
// func (c *cRoleaccount) DeleteRoleAccount(ctx *gin.Context) {
// 	var ids []string
// 	if err := ctx.ShouldBindJSON(&ids); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
// 		return
// 	}
// 	for _, id := range ids {
// 		if _, err := uuid.Parse(id); err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid id: %s", id)})
// 			return
// 		}
// 	}
// 	code, err := service.RoleAccountItem().DeleteRoleAccount(ctx.Request.Context(), ids)
// 	if err != nil {
// 		response.ErrorResponse(ctx, code, err.Error())
// 		return
// 	}
// 	response.SuccessResponse(ctx, code, nil)
// }

// // UpdateRoleAccount
// // @Summary      Cập nhật Role account
// // @Description  API này cập nhật thông tin role account dựa trên ID
// // @Tags         RoleAccount
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Security     ApiKeyAuth
// // @Param        id   path   string  true  "ID role account cần cập nhật"
// // @Param        body body   model.RoleAccount true "Dữ liệu cập nhật role account"
// // @Success      200  {object}  response.ResponseData
// // @Failure      400  {object}  response.ErrorResponseData
// // @Failure      500  {object}  response.ErrorResponseData
// // @Router       /roleaccount/update_role_account/{id} [PUT]
// func (c *cRoleaccount) UpdateRoleAccount(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	var modelRoleAccount model.RoleAccount
// 	if err := ctx.ShouldBindJSON(&modelRoleAccount); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
// 		return
// 	}
// 	if _, err := uuid.Parse(id); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role id"})
// 		return
// 	}
// 	code, roleAccount, err := service.RoleAccountItem().UpdateRoleAccount(ctx.Request.Context(), id, &modelRoleAccount)
// 	if err != nil {
// 		response.ErrorResponse(ctx, code, err.Error())
// 		return
// 	}
// 	response.SuccessResponse(ctx, code, roleAccount)
// }
