package service

import (
	"context"
	"go-backend-api/internal/model"
)

type (
	Iuser interface {
		CreateUser(ctx context.Context, in *model.AccountInput) (codeResult int, out model.AccountOutput, err error)
	}
)

var (
	localUser Iuser
)

func UserItem() Iuser {
	if localUser == nil {
		panic("implement localUserLogin not found for interface IUserLogin")
	}
	return localUser
}

func InitUserItem(i Iuser) {
	localUser = i
}
