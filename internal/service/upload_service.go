package service

import "mime/multipart"

type (
	IUpload interface {
		UploadFile(file multipart.File, header *multipart.FileHeader) (string, error)
	}
)

var (
	localUpload IUpload
)

func UploadItem() IUpload {
	if localUpload == nil {
		panic("implement localUpload not found for interface IUpload")
	}
	return localUpload
}

func InitUploadItem(i IUpload) {
	localUpload = i
}
