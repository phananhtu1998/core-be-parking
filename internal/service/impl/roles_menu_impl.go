package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"
	"time"

	"github.com/google/uuid"
)

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
	roleMenus, err := s.r.GetRoleMenuByRoleId(ctx, database.GetRoleMenuByRoleIdParams{
		ID:       roleId,
		RoleName: "%" + search + "%",
	})
	if err != nil {
		return response.ErrCodeRoleMenuError, nil, err
	}

	result := make([]model.RoleMenuOutput, 0, len(roleMenus))
	for _, rm := range roleMenus {
		var methods []string
		if len(rm.ListMethod) > 0 {
			if err := json.Unmarshal(rm.ListMethod, &methods); err != nil {
				return response.ErrCodeRoleMenuError, nil, err
			}
		}
		result = append(result, model.RoleMenuOutput{
			Id:              rm.ID,
			Menu_name:       rm.MenuName,
			Menu_url:        rm.MenuUrl,
			Menu_icon:       rm.MenuIcon,
			Menu_group_name: rm.MenuGroupName,
			Role_code:       rm.Code,
			Role_name:       rm.RoleName,
			RolesMenu: model.RolesMenu{
				Menu_id:    rm.MenuID,
				Role_id:    rm.RoleID,
				ListMethod: methods,
			},
		})
	}

	return response.ErrCodeRoleMenuSucces, result, nil
}

func (s *sRolesMenu) UpdateRolesMenu(ctx context.Context, id string, in *model.RolesMenu) (int, model.RolesMenu, error) {
	// Convert []string to json.RawMessage for database storage
	listMethodJSON, err := json.Marshal(in.ListMethod)
	if err != nil {
		return response.ErrCodeRoleMenuError, model.RolesMenu{}, err
	}

	// Thực hiện update vào database
	err = s.r.UpdateRolesMenu(ctx, database.UpdateRolesMenuParams{
		ID:         id,
		MenuID:     in.Menu_id,
		RoleID:     in.Role_id,
		ListMethod: listMethodJSON,
	})

	if err != nil {
		return response.ErrCodeRoleMenuError, model.RolesMenu{}, err
	}

	// Trả về dữ liệu đã update thành công
	return response.ErrCodeRoleMenuSucces, *in, nil
}

func (s *sRolesMenu) DeleteRolesMenu(ctx context.Context, id string) (int, error) {
	err := s.r.DeleteRolesMenu(ctx, database.DeleteRolesMenuParams{
		ID:       id,
		UpdateAt: time.Now(),
	})
	if err != nil {
		return response.ErrCodeRoleMenuError, err
	}

	return response.ErrCodeRoleMenuSucces, nil
}

func (s *sRolesMenu) CreateMultipleRoleMenus(ctx context.Context, inputs []model.RolesMenu) (codeResult int, out []model.RolesMenu, err error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return response.ErrCodeMenuErrror, nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	var committed bool
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()
	outputs := make([]model.RolesMenu, 0, len(inputs))
	for _, input := range inputs {
		newUUID := uuid.New().String()
		listMethodJSON, err := json.Marshal(input.ListMethod)
		if err != nil {
			return response.ErrCodeRoleMenuError, out, err
		}
		if err := s.r.CreateRolesMenu(ctx, database.CreateRolesMenuParams{
			ID:         newUUID,
			MenuID:     input.Menu_id,
			RoleID:     input.Role_id,
			ListMethod: listMethodJSON,
		}); err != nil {
			return response.ErrCodeMenuErrror, nil, err
		}
		outputs = append(outputs, model.RolesMenu{
			Menu_id:    input.Menu_id,
			Role_id:    input.Role_id,
			ListMethod: input.ListMethod,
		})
	}
	if err = tx.Commit(); err != nil {
		return response.ErrCodeMenuErrror, nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	committed = true
	return response.ErrCodeSucces, outputs, nil
}
