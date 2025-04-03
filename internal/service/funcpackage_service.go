package service

import (
	"context"
	"go-backend-api/internal/model"
)

type (
	Ifuncpackage interface {
		CreateFuncPackage(ctx context.Context, in *model.FuncpackageInput) (codeResult int, out model.FuncpackageOutput, err error)
		GetAlFuncPackageByCreatedBy(ctx context.Context) (codeResult int, out []model.FuncpackageOutput, err error)
	}
)

var (
	localFuncpackage Ifuncpackage
)

func FuncpackageItem() Ifuncpackage {
	if localFuncpackage == nil {
		panic("implement localFuncpackage not found for interface IFuncpackage")
	}
	return localFuncpackage
}

func InitFuncpackageItem(i Ifuncpackage) {
	localFuncpackage = i
}
