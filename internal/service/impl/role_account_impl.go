package impl

import (
	"context"
	"go-backend-api/internal/database"
	"go-backend-api/internal/model"
	"go-backend-api/pkg/response"
	"time"

	"github.com/google/uuid"
)

type sRoleAccount struct {
	r *database.Queries
}

func NewRoleAccountImpl(r *database.Queries) *sRoleAccount {
	return &sRoleAccount{r: r}
}

func (s *sRoleAccount) CreateRoleAccount(ctx context.Context, roleAccount *model.RoleAccount) (codeResult int, out model.RoleAccountOutput, err error) {
	Id := uuid.New().String()
	err = s.r.CreateRoleAccount(ctx, database.CreateRoleAccountParams{
		ID:        Id,
		RoleID:    roleAccount.Role_id,
		AccountID: roleAccount.Account_id,
		LicenseID: roleAccount.License_id,
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
				License_id: item.LicenseID,
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
				License_id: item.LicenseID,
			},
			Create_at: item.CreateAt,
			Update_at: item.UpdateAt.Format("02-01-2006 15:04:05"),
		})
	}
	return codeResult, out, err
}
