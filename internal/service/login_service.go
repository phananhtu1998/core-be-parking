package service

import (
	"context"
	"go-backend-api/internal/model"
)

type (
	ILogin interface {
		Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error)
		Logout(ctx context.Context) (codeResult int, err error)
		RefreshTokens(ctx context.Context) (codeResult int, out model.LoginOutput, err error)
		ChangePassword(ctx context.Context, in *model.ChangePasswordInput) (codeResult int, out model.LoginOutput, err error)
	}
)

var (
	localLogin ILogin
)

func LoginItem() ILogin {
	if localLogin == nil {
		panic("implement localUserLogin not found for interface ILogin")
	}
	return localLogin
}

func InitLoginItem(i ILogin) {
	localLogin = i
}
