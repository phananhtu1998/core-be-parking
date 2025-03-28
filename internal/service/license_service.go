package service

import (
	"context"
	"go-backend-api/internal/model"
)

type (
	ILicense interface {
		CreateLicense(ctx context.Context, in *model.License) (codeResult int, out model.LicenseOutput, err error)
	}
)

var (
	localLicense ILicense
)

func LicenseItem() ILicense {
	if localLicense == nil {
		panic("implement localUserLicense not found for interface ILicense")
	}
	return localLicense
}

func InitLicenseItem(i ILicense) {
	localLicense = i
}
