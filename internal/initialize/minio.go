package initialize

import (
	"context"
	"go-backend-api/global"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinio() error {
	minioClient, err := minio.New(global.Config.MinIO.ENDPOINT, &minio.Options{
		Creds:  credentials.NewStaticV4(global.Config.MinIO.ACCESS_KEY, global.Config.MinIO.SECRET_KEY, ""),
		Secure: global.Config.MinIO.USESSL,
	})
	if err != nil {
		log.Println("Lỗi kết nối với minio: ", err)
		return err
	}
	global.MinioClient = minioClient
	// Tạo bucket nếu chưa có
	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, global.Config.MinIO.BUCKET_NAME)
	if err != nil {
		log.Println("Lỗi kết nối với minio: ", err)
		return err
	}
	if !exists {
		err = minioClient.MakeBucket(ctx, global.Config.MinIO.BUCKET_NAME, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
		log.Println("Bucket đã được tạo:", global.Config.MinIO.BUCKET_NAME)
	}

	return nil
}
