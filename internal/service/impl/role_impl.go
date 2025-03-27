package impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/internal/utils"
	"go-backend-api/pkg/response"
	"time"

	"github.com/google/uuid"
)

type sRole struct {
	r   *database.Queries
	qTx *sql.Tx
	db  *sql.DB
}

func NewRoleImpl(r *database.Queries, qTx *sql.Tx, db *sql.DB) service.IRole {
	return &sRole{
		r:   r,
		qTx: qTx,
		db:  db,
	}
}

func (s *sRole) CreateRole(ctx context.Context, in *model.Role) (codeResult int, out model.Role, err error) {
	var leftValue, rightValue int32
	newID := uuid.New().String()

	// Bắt đầu transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Nếu created_by trống, đây là node gốc
	if in.Created_by == "" {
		// Lấy giá trị right lớn nhất
		maxRightValue, err := s.r.GetMaxRightValue(ctx)
		if err != nil {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to get max right value: %w", err)
		}

		// Đặt node mới là root - Sửa phần này
		maxRightValueInt64 := maxRightValue.(int64)
		leftValue = int32(maxRightValueInt64) + 1
		rightValue = int32(maxRightValueInt64) + 2
	} else {
		// Lấy thông tin của node cha
		parentRole, err := s.r.GetParentRoleInfo(ctx, in.Created_by)
		if err != nil {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("parent role not found: %w", err)
		}

		// Cập nhật right values
		err = s.r.UpdateRightValuesForInsert(ctx, parentRole.RoleRightValue)
		if err != nil {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to update right values: %w", err)
		}

		// Cập nhật left values
		err = s.r.UpdateLeftValuesForInsert(ctx, parentRole.RoleRightValue)
		if err != nil {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to update left values: %w", err)
		}

		// Đặt giá trị cho node mới
		leftValue = parentRole.RoleRightValue
		rightValue = parentRole.RoleRightValue + 1
	}

	// Tạo role mới
	_, err = s.r.CreateRole(ctx, database.CreateRoleParams{
		ID:             newID,
		Code:           in.Code,
		RoleName:       in.Role_name,
		RoleLeftValue:  leftValue,
		RoleRightValue: rightValue,
		RoleMaxNumber:  int64(in.Role_max_number),
		IsLicensed:     in.Is_licensed,
		CreatedBy:      in.Created_by,
	})
	if err != nil {
		return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to create role: %w", err)
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return response.ErrCodeSucces, model.Role{
		Id:               newID,
		Code:             in.Code,
		Role_name:        in.Role_name,
		Role_left_value:  int(leftValue),
		Role_right_value: int(rightValue),
		Role_max_number:  in.Role_max_number,
		Is_licensed:      in.Is_licensed,
		Created_by:       in.Created_by,
		Created_at:       time.Now(),
	}, nil
}

// GetAllRoles - Lấy danh sách vai trò phân trang + xây dựng cây hierarchy
func (s *sRole) GetAllRoles(ctx context.Context, page, pageSize int) (codeResult int, out []model.RoleHierarchyOutput, total int64, err error) {
	// Lấy tổng số records
	total, err = s.r.GetTotalRoles(ctx)
	if err != nil {
		return response.ErrCodeRoleError, nil, 0, err
	}

	// Tính offset
	offset := (page - 1) * pageSize
	roles, err := s.r.GetAllRole(ctx, database.GetAllRoleParams{
		Limit:  int32(pageSize),
		Offset: int32(offset),
	})
	if err != nil {
		return response.ErrCodeRoleError, nil, 0, err
	}

	// Chuyển đổi danh sách Role sang Model
	var modelRoles []model.Role
	for _, r := range roles {
		modelRoles = append(modelRoles, utils.ConvertToModelRole(r))
	}

	// Thay đổi kiểu của map từ map[int]bool thành map[string]bool
	childMap := make(map[string]bool)
	for _, role := range roles {
		for _, parent := range roles {
			if role.ID != parent.ID &&
				role.RoleLeftValue > parent.RoleLeftValue &&
				role.RoleRightValue < parent.RoleRightValue {
				childMap[role.ID] = true
				break
			}
		}
	}

	// Xây dựng danh sách root nodes
	var rootNodes []model.RoleHierarchyOutput
	for _, role := range roles {
		if !childMap[role.ID] { // Chỉ lấy các role không phải con của role khác
			node := utils.BuildRoleHierarchy(utils.ConvertToModelRole(role), modelRoles)
			rootNodes = append(rootNodes, node)
		}
	}

	return response.ErrCodeSucces, rootNodes, total, nil
}

func (s *sRole) GetRoleById(ctx context.Context, parentId string) (codeResult int, out []model.RoleHierarchyOutput, err error) {
	// First, get the role by its ID
	parentRole, err := s.r.GetRoleById(ctx, parentId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeRoleNotFound, nil, fmt.Errorf("role not found")
		}
		return response.ErrCodeRoleError, nil, fmt.Errorf("failed to get role: %w", err)
	}

	// Get all child roles where created_by matches the roleId
	childRoles, err := s.r.GetChildRolesByParentId(ctx, parentId)
	if err != nil {
		return response.ErrCodeRoleError, nil, fmt.Errorf("failed to get child roles: %w", err)
	}

	// Create the parent node
	rootNode := model.RoleHierarchyOutput{
		Id:        parentRole.ID,
		Code:      parentRole.Code,
		Role_name: parentRole.RoleName,
		Children:  make([]model.RoleHierarchyOutput, 0, len(childRoles)),
	}

	// Add child nodes
	for _, role := range childRoles {
		childNode := model.RoleHierarchyOutput{
			Id:        role.ID,
			Code:      role.Code,
			Role_name: role.RoleName,
			Children:  []model.RoleHierarchyOutput{}, // Initialize empty children array
		}
		rootNode.Children = append(rootNode.Children, childNode)
	}

	// Create a slice with the root node as the only element
	result := []model.RoleHierarchyOutput{rootNode}

	return response.ErrCodeSucces, result, nil
}

func (s *sRole) DeleteRole(ctx context.Context, id string) (codeResult int, err error) {
	// Bắt đầu transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return response.ErrCodeRoleError, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Lấy role_left_value và role_right_value của role cha
	parentRole, err := s.r.GetParentRoleInfo(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeRoleNotFound, fmt.Errorf("role not found")
		}
		return response.ErrCodeRoleError, fmt.Errorf("failed to get parent role info: %w", err)
	}

	// Xóa mềm role cha và tất cả role con trong một truy vấn duy nhất
	err = s.r.SoftDeleteRolesByRange(ctx, database.SoftDeleteRolesByRangeParams{
		UpdateAt:       time.Now(),
		RoleLeftValue:  parentRole.RoleLeftValue,
		RoleRightValue: parentRole.RoleRightValue,
	})
	if err != nil {
		return response.ErrCodeRoleError, fmt.Errorf("failed to delete roles: %w", err)
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return response.ErrCodeRoleError, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return response.ErrCodeSucces, nil
}

func (s *sRole) GetAllPermission(ctx context.Context) (codeResult int, out []model.RolePermission, err error) {
	// Lấy tất cả quyền của vai trò
	permissions, err := s.r.GetAllPermissions(ctx)
	if err != nil {
		return response.ErrCodeRoleError, nil, fmt.Errorf("failed to get all permissions: %w", err)
	}
	for _, item := range permissions {
		out = append(out, model.RolePermission{
			Id:              item.ID,
			Role_name:       item.RoleName,
			Menu_group_name: item.MenuGroupName,
			Method:          string(item.Method),
		})
	}
	return response.ErrCodeSucces, out, nil
}
