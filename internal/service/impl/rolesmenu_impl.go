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
