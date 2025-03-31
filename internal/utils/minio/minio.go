package minio

import (
	"context"
	"fmt"
	"go-backend-api/global"
	"go-backend-api/internal/initialize"
	"mime/multipart"
	"time"

	"github.com/minio/minio-go/v7"
)

// UploadChunkFileToMinio tải file lên MinIO
func UploadChunkFileToMinio(file *multipart.FileHeader) (string, error) {
	// Lấy client từ hàm khởi tạo
	client, err := initialize.InitMinio()
	if err != nil {
		return "", fmt.Errorf("không thể kết nối với MinIO: %w", err)
	}

	// Mở file từ request
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("không thể mở file: %w", err)
	}
	defer src.Close()

	// Tạo tên đối tượng duy nhất bằng timestamp và tên file
	objectName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	// Xác định Content-Type
	contentType := file.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	// Upload file
	_, err = client.PutObject(
		ctx,
		global.Config.MinIO.BUCKET_NAME,
		objectName,
		src,
		file.Size,
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)
	if err != nil {
		return "", fmt.Errorf("lỗi khi tải lên file: %w", err)
	}

	// Trả về URL của file đã tải lên
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


