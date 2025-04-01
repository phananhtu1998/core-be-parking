package impl

import (
	"context"
	"database/sql"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
)

type sFuncpackage struct {
	r   *database.Queries
	qTx *sql.Tx
	db  *sql.DB
}

func NewFuncpackageImpl(r *database.Queries, qTx *sql.Tx, db *sql.DB) service.Ifuncpackage {
	return &sFuncpackage{
		r:   r,
		qTx: qTx,
		db:  db,
	}
}

func (s *sFuncpackage) CreateFuncPackage(ctx context.Context, in *model.Role) (codeResult int, out model.Role, err error) {
	return codeResult, out, err
}
