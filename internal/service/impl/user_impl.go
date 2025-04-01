package impl

import (
	"context"
	"database/sql"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
)

type sUser struct {
	r   *database.Queries
	qTx *sql.Tx
	db  *sql.DB
}

func NewUserImpl(r *database.Queries, qTx *sql.Tx, db *sql.DB) *sUser {
	return &sUser{
		r:   r,
		qTx: qTx,
		db:  db,
	}
}

func (s *sUser) CreateUser(ctx context.Context, in *model.AccountInput) (codeResult int, out model.AccountOutput, err error) {
	return
}
