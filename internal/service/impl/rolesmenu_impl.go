package impl

import (
	"context"
	"database/sql"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/pkg/response"
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
	// Thực hiện insert vào database
	err = s.r.CreateRolesMenu(ctx, database.CreateRolesMenuParams{
		ID:         in.Id,
		MenuID:     in.Menu_id,
		RoleID:     in.Role_id,
		ListMethod: in.ListMethod,
	})

	if err != nil {
		return response.ErrCodeRoleMenuError, model.RolesMenu{}, err
	}

	// Trả về dữ liệu đã insert thành công
	return response.ErrCodeRoleMenuSucces, *in, nil
}
