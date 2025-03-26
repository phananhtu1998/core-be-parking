package service

import (
	"context"
	"go-backend-api/internal/model"
)

type (
	IRoleAccount interface {
		CreateRoleAccount(ctx context.Context, roleAccount *model.RoleAccount) (codeResult int, out model.RoleAccountOutput, err error)
		GetAllRoleAccountByRoleId(ctx context.Context, role_Id string) (codeResult int, out []model.RoleAccountOutput, err error)
		GetAllRoleAccountByAccountId(ctx context.Context, account_Id string) (codeResult int, out []model.RoleAccountOutput, err error)
		UpdateRoleAccount(ctx context.Context, id string, roleAccount *model.RoleAccount) (codeResult int, out model.RoleAccountOutput, err error)
		DeleteRoleAccount(ctx context.Context, id []string) (codeResult int, err error)
	}
)

var (
	localRoleAccount IRoleAccount
)

func RoleAccountItem() IRoleAccount {
	if localRoleAccount == nil {
		panic("implement localRoleAccount not found for interface IRoleAccount")
	}
	return localRoleAccount
}

func InitRoleAccountItem(i IRoleAccount) {
	localRoleAccount = i
}
