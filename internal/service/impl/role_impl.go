package impl

import (
	"context"
	"database/sql"
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

// GetAllRoles retrieves all roles and builds a hierarchical tree
func (s *sRole) GetAllRoles(ctx context.Context) (codeResult int, out []model.RoleHierarchyOutput, err error) {
	roles, err := s.r.GetAllRole(ctx)
	if err != nil {
		return response.ErrCodeRoleError, nil, err
	}

	var modelRoles []model.Role
	for _, r := range roles {
		modelRoles = append(modelRoles, utils.ConvertToModelRole(r))
	}

	var rootNodes []model.RoleHierarchyOutput
	for _, role := range roles {
		isChild := false
		for _, potentialParent := range roles {
			if role.ID != potentialParent.ID &&
				role.RoleLeftValue > potentialParent.RoleLeftValue &&
				role.RoleRightValue < potentialParent.RoleRightValue {
				isChild = true
				break
			}
		}

		if !isChild {
			node := utils.BuildRoleHierarchy(utils.ConvertToModelRole(role), modelRoles)
			rootNodes = append(rootNodes, node)
		}
	}

	return response.ErrCodeSucces, rootNodes, nil
}
