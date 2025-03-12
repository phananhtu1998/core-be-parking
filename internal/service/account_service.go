package service

import (
	"context"
	"go-backend-api/internal/model"
)

type (
	Iaccount interface {
		CreateAccount(ctx context.Context, in *model.AccountInput) (codeResult int, out model.AccountOutput, err error)
		GetAccountById(ctx context.Context, id string) (codeResult int, out model.AccountOutput, err error)
		UpdateAccount(ctx context.Context, in *model.AccountInput, id string) (codeResult int, out model.AccountOutput, err error)
		DeleteAccount(ctx context.Context, id string) (codeResult int, err error)
		GetAllAccount(ctx context.Context) (codeResult int, out []model.AccountOutput, err error)
	}
)

var (
	localAccount Iaccount
)

func AccountItem() Iaccount {
	if localAccount == nil {
		panic("implement localUserLogin not found for interface IUserLogin")
	}
	return localAccount
}

func InitAccountItem(i Iaccount) {
	localAccount = i
}
