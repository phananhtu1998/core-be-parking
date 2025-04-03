package role

// import (
// 	"go-backend-api/internal/model"
// 	"go-backend-api/internal/service"
// 	"go-backend-api/pkg/response"
// 	"log"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// var Roles = new(cRole)

// type cRole struct {
// 	roleService service.IRole
// }

// // Role
// // @Summary Tạo vai trò
// // @Description Api tạo vai trò trong hệ thống
// // @Tags Role
// // @Accept json
// // @Produce json
// // @Security     BearerAuth
// // @Security     ApiKeyAuth
// // @Param        payload body model.Role true "payload"
// // @Success 200 {object} response.ResponseData
// // @Failure      500  {object}  response.ErrorResponseData
// // @Router /role/create_role [post]
// func (c *cRole) CreateRole(ctx *gin.Context) {
// 	var params model.Role
// 	if err := ctx.ShouldBindJSON(&params); err != nil {
// 		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
// 		return
// 	}
// 	codeRole, dataRole, err := service.RoleItem().CreateRole(ctx.Request.Context(), &params)
// 	if err != nil {
// 		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
// 		return
// 	}
// 	response.SuccessResponse(ctx, codeRole, dataRole)
// }
// // GetAllRoles
// // @Summary Lấy danh sách role
// // @Description Api lấy danh sách role trong hệ thống
// // @Tags Role
// // @Accept json
// // @Produce json
// // @Security     BearerAuth
// // @Security     ApiKeyAuth
// // @Param page query int false "Số trang (mặc định: 1)"
// // @Param page_size query int false "Số lượng mỗi trang (mặc định: 20)"
// // @Success 200 {object} response.ResponseData
// // @Failure      500  {object}  response.ErrorResponseData
// // @Router /role/get_all_roles [get]
// func (c *cRole) GetAllRoles(ctx *gin.Context) {
// 	// Lấy tham số phân trang từ query params
// 	page := 1
// 	pageSize := 20

// 	if pageStr := ctx.Query("page"); pageStr != "" {
// 		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
// 			page = p
// 		}
// 	}

// 	if pageSizeStr := ctx.Query("page_size"); pageSizeStr != "" {
// 		if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 {
// 			pageSize = ps
// 		}
// 	}
// 	codeRole, dataRole, total, err := service.RoleItem().GetAllRoles(ctx.Request.Context(), page, pageSize)
// 	if err != nil {
// 		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
// 		return
// 	}
// 	// Trả về response với thông tin phân trang
// 	paginatedResponse := model.NewPaginationResponse(dataRole, page, pageSize, total)
// 	response.SuccessResponse(ctx, codeRole, paginatedResponse)
// }

// // GetRoleById
// // @Summary      Lấy role theo ID
// // @Description  API này trả về role theo ID
// // @Tags         Role
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Security     ApiKeyAuth
// // @Param        id   path   string  true  "ID role cần lấy"
// // @Success      200  {object}  response.ResponseData
// // @Failure      500  {object}  response.ErrorResponseData
// // @Router       /role/get_role_by_id/{id} [GET]
// func (c *cRole) GetRoleById(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	code, menu, err := service.RoleItem().GetRoleById(ctx.Request.Context(), id)
// 	if err != nil {
// 		log.Printf("Error getting menu: %v", err)
// 		response.ErrorResponse(ctx, code, err.Error())
// 		return
// 	}

// 	response.SuccessResponse(ctx, code, menu)
// }

// // DeleteRole
// // @Summary Xóa role
// // @Description API xóa mềm role trong hệ thống
// // @Tags Role
// // @Accept json
// // @Produce json
// // @Security     BearerAuth
// // @Security     ApiKeyAuth
// // @Param        id   path      string  true  "ID của role cần xóa"
// // @Success      200  {object}  response.ResponseData
// // @Failure      500  {object}  response.ErrorResponseData
// // @Router       /role/delete_role/{id} [delete]
// func (c *cRole) DeleteRole(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	code, err := service.RoleItem().DeleteRole(ctx.Request.Context(), id)
// 	if err != nil {
// 		response.ErrorResponse(ctx, code, err.Error())
// 		return
// 	}
// 	response.SuccessResponse(ctx, code, nil)
// }
