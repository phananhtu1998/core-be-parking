package utils

import (
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
)

// Convert database row to model role
func ConvertToModelRole(dbRole database.GetAllRoleRow) model.Role {
	return model.Role{
		Id:               dbRole.ID,
		Code:             dbRole.Code,
		Role_name:        dbRole.RoleName,
		Role_left_value:  int(dbRole.RoleLeftValue),
		Role_right_value: int(dbRole.RoleRightValue),
		Role_max_number:  int(dbRole.RoleMaxNumber),
		Created_by:       dbRole.CreatedBy,
		Created_at:       dbRole.CreateAt,
	}
}

// Helper function to build hierarchy recursively
func BuildRoleHierarchy(role model.Role, allRoles []model.Role) model.RoleHierarchyOutput {
	node := model.RoleHierarchyOutput{
		Id:        role.Id,
		Code:      role.Code,
		Role_name: role.Role_name,
		Children:  []model.RoleHierarchyOutput{},
	}

	// Find immediate children
	for _, childRole := range allRoles {
		if childRole.Id != role.Id &&
			childRole.Role_left_value > role.Role_left_value &&
			childRole.Role_right_value < role.Role_right_value {

			hasIntermediateParent := false
			for _, potentialIntermediate := range allRoles {
				if potentialIntermediate.Id != role.Id &&
					potentialIntermediate.Id != childRole.Id &&
					childRole.Role_left_value > potentialIntermediate.Role_left_value &&
					childRole.Role_right_value < potentialIntermediate.Role_right_value &&
					potentialIntermediate.Role_left_value > role.Role_left_value &&
					potentialIntermediate.Role_right_value < role.Role_right_value {
					hasIntermediateParent = true
					break
				}
			}

			if !hasIntermediateParent {
				childNode := BuildRoleHierarchy(childRole, allRoles)
				node.Children = append(node.Children, childNode)
			}
		}
	}

	return node
}
