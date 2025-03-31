package impl

import (
	"context"
	"fmt"
	"go-backend-api/global"
	"mime/multipart"
	"time"

	"github.com/minio/minio-go/v7"
)

type sUpload struct{}

func NewUploadImpl() *sUpload {
	return &sUpload{}
}
func (s *sUpload) UploadFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	client := global.MinioClient // Sử dụng client từ global để tránh import cycle

	if client == nil {
		return "", fmt.Errorf("MinIO client chưa được khởi tạo")
	}

	// Tạo tên file duy nhất
	objectName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), header.Filename)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	// Xác định Content-Type
	contentType := header.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	// Upload file lên MinIO
	_, err := client.PutObject(
		ctx,
		global.Config.MinIO.BUCKET_NAME,
		objectName,
		file,
		header.Size,
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		return "", fmt.Errorf("Lỗi khi tải lên file: %w", err)
	}

	// Trả về URL file
	protocol := "http"
	if global.Config.MinIO.USESSL {
		protocol = "https"
	}

	fileURL := fmt.Sprintf("%s://%s/%s/%s",
		protocol,
		global.Config.MinIO.ENDPOINT,
		global.Config.MinIO.BUCKET_NAME,
		objectName,
	)

	return fileURL, nil
}
