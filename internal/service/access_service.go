package service

import (
	"context"
	"go-backend-api/internal/model"
)

type (
	IAccess interface {
		CreateAccount(ctx context.Context, in *model.AccountInput) (codeResult int, err error)
	}
)

var (
	localAccess IAccess
)

func AccessItem() IAccess {
	if localAccess == nil {
		panic("implement access key")
	}
	return localAccess
}

func InitAccessItem(i IAccess) {
	localAccess = i
}
