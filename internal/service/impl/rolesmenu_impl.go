package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"

	"github.com/google/uuid"
)

const getRoleMenuByRoleId = `
SELECT 
    m.id, 
    m.menu_name, 
    m.menu_url, 
    m.menu_icon, 
    rm.menu_id,
    rm.role_id,
    m.menu_group_name, 
    r.code, 
    r.role_name, 
    rm.list_method 
FROM roles_menu rm
JOIN menu m ON m.id = rm.menu_id AND m.is_deleted = FALSE
JOIN role r ON r.id = rm.role_id AND r.is_deleted = FALSE
WHERE r.id = ?
AND (
    ? = '' OR MATCH(r.role_name) AGAINST (? IN NATURAL LANGUAGE MODE)
)`

type sRolesMenu struct {
	r   *database.Queries
	qTx *sql.Tx
	db  *sql.DB
}

func NewRolesMenuImpl(r *database.Queries, qTx *sql.Tx, db *sql.DB) service.IRolesMenu {
	return &sRolesMenu{
		r:   r,
		qTx: qTx,
		db:  db,
	}
}

func (s *sRolesMenu) CreateRolesMenu(ctx context.Context, in *model.RolesMenu) (codeResult int, out model.RolesMenu, err error) {
	// Convert []string to json.RawMessage for database storage
	listMethodJSON, err := json.Marshal(in.ListMethod)
	if err != nil {
		return response.ErrCodeRoleMenuError, model.RolesMenu{}, err
	}

	// Thực hiện insert vào database
	err = s.r.CreateRolesMenu(ctx, database.CreateRolesMenuParams{
		ID:         uuid.New().String(),
		MenuID:     in.Menu_id,
		RoleID:     in.Role_id,
		ListMethod: listMethodJSON,
	})

	if err != nil {
		return response.ErrCodeRoleMenuError, model.RolesMenu{}, err
	}

	// Trả về dữ liệu đã insert thành công
	return response.ErrCodeRoleMenuSucces, *in, nil
}

func (s *sRolesMenu) GetRoleMenuByRoleId(ctx context.Context, roleId, search string) (int, []model.RoleMenuOutput, error) {
	stmt, err := s.db.PrepareContext(ctx, getRoleMenuByRoleId)
	if err != nil {
		return response.ErrCodeRoleMenuError, nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, roleId, search, search)
	if err != nil {
		return response.ErrCodeRoleMenuError, nil, err
	}
	defer rows.Close()

	var roleMenus []database.GetRoleMenuByRoleIdRow
	for rows.Next() {
		var rm database.GetRoleMenuByRoleIdRow
		if err := rows.Scan(
			&rm.ID,
			&rm.MenuName,
			&rm.MenuUrl,
			&rm.MenuIcon,
			&rm.MenuID,
			&rm.RoleID,
			&rm.MenuGroupName,
			&rm.Code,
			&rm.RoleName,
			&rm.ListMethod,
		); err != nil {
			return response.ErrCodeRoleMenuError, nil, err
		}
		roleMenus = append(roleMenus, rm)
	}

	result := make([]model.RoleMenuOutput, 0, len(roleMenus))
	for _, roleMenu := range roleMenus {
		var methods []string
		if len(roleMenu.ListMethod) > 0 {
			if err := json.Unmarshal(roleMenu.ListMethod, &methods); err != nil {
				return response.ErrCodeRoleMenuError, nil, err
			}
		}

		result = append(result, model.RoleMenuOutput{
			Id:              roleMenu.ID,
			Menu_name:       roleMenu.MenuName,
			Menu_url:        roleMenu.MenuUrl,
			Menu_icon:       roleMenu.MenuIcon,
			Menu_group_name: roleMenu.MenuGroupName,
			Role_code:       roleMenu.Code,
			Role_name:       roleMenu.RoleName,
			RolesMenu: model.RolesMenu{
				Menu_id:    roleMenu.MenuID,
				Role_id:    roleMenu.RoleID,
				ListMethod: methods,
			},
		})
	}

	return response.ErrCodeRoleMenuSucces, result, nil
}
