package service

import (
	"context"
	"go-backend-api/internal/model"
)

type (
	IMenu interface {
		CreateMenu(ctx context.Context, in *model.MenuInput) (codeResult int, out model.MenuOutput, err error)
	}
)

var (
	localMenu IMenu
)

func MenuItem() IMenu {
	if localMenu == nil {
		panic("implement localMenu not found for interface IMenu")
	}
	return localMenu
}

func InitMenuItem(i IMenu) {
	localMenu = i
}
