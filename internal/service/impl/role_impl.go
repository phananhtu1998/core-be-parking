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
}

func NewRoleImpl(r *database.Queries, qTx *sql.Tx) service.IRole {
	return &sRole{
		r:   r,
		qTx: qTx,
	}
}

// CreateRole creates a new role with binary tree structure
func (s *sRole) CreateRole(ctx context.Context, in *model.Role) (codeResult int, out model.Role, err error) {
	var leftValue, rightValue int32

	// If created_by is empty, this is a root role
	if in.Created_by == "" {
		// Get the maximum right_value
		roles, err := s.r.GetAllRole(ctx)
		if err != nil {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to get roles: %w", err)
		}

		var maxRightValue int32
		for _, role := range roles {
			if role.RoleRightValue > maxRightValue {
				maxRightValue = role.RoleRightValue
			}
		}

		// Set new role as root
		leftValue = maxRightValue + 1
		rightValue = maxRightValue + 2
	} else {
		// Get parent role info
		parentRole, err := s.r.GetParentRoleInfo(ctx, in.Created_by)
		if err != nil {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("parent role not found: %w", err)
		}

		// Update existing roles to make space for new role
		_, err = s.r.UpdateRoleTree(ctx, parentRole.RoleRightValue)
		if err != nil {
			return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to update role tree: %w", err)
		}

		// Set new role values
		leftValue = parentRole.RoleRightValue
		rightValue = parentRole.RoleRightValue + 1
	}

	// Create new role
	role, err := s.r.CreateRole(ctx, database.CreateRoleParams{
		ID:             uuid.New().String(),
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

	// Get the inserted ID
	id, err := role.LastInsertId()
	if err != nil {
		return response.ErrCodeRoleError, model.Role{}, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return response.ErrCodeSucces, model.Role{
		Id:               fmt.Sprintf("%d", id),
		Code:             in.Code,
		Role_name:        in.Role_name,
		Role_left_value:  int(leftValue),
		Role_right_value: int(rightValue),
		Role_max_number:  in.Role_max_number,
		Is_licensed:      in.Is_licensed,
		Created_by:       in.Created_by,
		Created_at:       time.Now(),
	}, err
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
