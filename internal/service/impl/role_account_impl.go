package impl

import (
	"context"
	"database/sql"
	"fmt"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
	"go-backend-api/pkg/response"
	"time"

	"github.com/google/uuid"
)

type sRoleAccount struct {
	r   *database.Queries
	qTx *sql.Tx
	db  *sql.DB
}

func NewRoleAccountImpl(r *database.Queries, qTx *sql.Tx, db *sql.DB) *sRoleAccount {
	return &sRoleAccount{
		r:   r,
		qTx: qTx,
		db:  db,
	}
}

func (s *sRoleAccount) CreateRoleAccount(ctx context.Context, roleAccount *model.RoleAccount) (codeResult int, out model.RoleAccountOutput, err error) {
	Id := uuid.New().String()
	err = s.r.CreateRoleAccount(ctx, database.CreateRoleAccountParams{
		ID:        Id,
		RoleID:    roleAccount.Role_id,
		AccountID: roleAccount.Account_id,
	})
	roleaccount := model.RoleAccountOutput{
		Id: Id,
		RoleAccount: model.RoleAccount{
			Role_id:    roleAccount.Role_id,
			Account_id: roleAccount.Account_id,
			License_id: roleAccount.License_id,
		},
		Create_at: time.Now(),
		Update_at: string(time.Now().Format("02-01-2006 15:04:05")),
	}
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	return response.ErrCodeSucces, roleaccount, err
}
func (s *sRoleAccount) GetAllRoleAccountByRoleId(ctx context.Context, role_Id string) (codeResult int, out []model.RoleAccountOutput, err error) {
	roleAccount, err := s.r.GetRoleAccountByRoleId(ctx, role_Id)
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	for _, item := range roleAccount {
		out = append(out, model.RoleAccountOutput{
			Id: item.ID,
			RoleAccount: model.RoleAccount{
				Role_id:    item.RoleID,
				Account_id: item.AccountID,
			},
			Create_at: item.CreateAt,
			Update_at: item.UpdateAt.Format("02-01-2006 15:04:05"),
		})
	}
	return response.ErrCodeSucces, out, err
}
func (s *sRoleAccount) GetAllRoleAccountByAccountId(ctx context.Context, account_Id string) (codeResult int, out []model.RoleAccountOutput, err error) {
	roleAccount, err := s.r.GetRoleAccountByAccountId(ctx, account_Id)
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	for _, item := range roleAccount {
		out = append(out, model.RoleAccountOutput{
			Id: item.ID,
			RoleAccount: model.RoleAccount{
				Role_id:    item.RoleID,
				Account_id: item.AccountID,
			},
			Create_at: item.CreateAt,
			Update_at: item.UpdateAt.Format("02-01-2006 15:04:05"),
		})
	}
	return codeResult, out, err
}

func (s *sRoleAccount) DeleteRoleAccount(ctx context.Context, ids []string) (codeResult int, err error) {
	// Bắt đầu transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return response.ErrCodeRoleAccountError, fmt.Errorf("failed to begin transaction: %w", err)
	}
	// Đảm bảo rollback nếu có lỗi
	var committed bool
	defer func() {
		if !committed {
			tx.Rollback() // Gọi rollback nếu chưa commit
		}
	}()
	//
	for _, id := range ids {
		err := s.r.DeleteRoleAccount(ctx, id)
		if err != nil {
			return response.ErrCodeRoleAccountError, fmt.Errorf("failed to delete role account: %w", err)
		}
	}
	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return response.ErrCodeRoleAccountError, fmt.Errorf("failed to commit transaction: %w", err)
	}
	// Đánh dấu transaction đã commit
	committed = true
	return response.ErrCodeRoleAccountSucces, nil
}

func (s *sRoleAccount) UpdateRoleAccount(ctx context.Context, id string, roleAccount *model.RoleAccount) (codeResult int, out model.RoleAccountOutput, err error) {
	err = s.r.UpdateRoleAccount(ctx, database.UpdateRoleAccountParams{
		ID:        id,
		RoleID:    roleAccount.Role_id,
		AccountID: roleAccount.Account_id,
	})
	roleaccount := model.RoleAccountOutput{
		Id: id,
		RoleAccount: model.RoleAccount{
			Role_id:    roleAccount.Role_id,
			Account_id: roleAccount.Account_id,
			License_id: roleAccount.License_id,
		},
		Create_at: time.Now(),
		Update_at: string(time.Now().Format("02-01-2006 15:04:05")),
	}
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	return response.ErrCodeSucces, roleaccount, err
}
