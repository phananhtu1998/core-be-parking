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

// // UploadLargeFileToMinio tải file lớn lên MinIO bằng phương pháp Multipart Upload
// // UploadLargeFileToMinio tải file lớn lên MinIO sử dụng multipart upload
// func UploadLargeFileToMinio(file *multipart.FileHeader) (string, error) {
// 	// Lấy client từ hàm khởi tạo
// 	client, err := initialize.InitMinio()
// 	if err != nil {
// 		return "", fmt.Errorf("không thể kết nối với MinIO: %w", err)
// 	}

// 	// Mở file từ request
// 	src, err := file.Open()
// 	if err != nil {
// 		return "", fmt.Errorf("không thể mở file: %w", err)
// 	}
// 	defer src.Close()

// 	// Tạo tên đối tượng duy nhất bằng timestamp và tên file
// 	objectName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
// 	defer cancel()

// 	// Xác định Content-Type
// 	contentType := file.Header.Get("Content-Type")
// 	if contentType == "" {
// 		contentType = "application/octet-stream"
// 	}
// 	// Khởi tạo Multipart Upload
// 	uploadID, err := client.NewMultipartUpload(ctx, global.Config.MinIO.BUCKET_NAME, objectName, minio.PutObjectOptions{
// 		ContentType: contentType,
// 	})
// 	if err != nil {
// 		return "", fmt.Errorf("lỗi khởi tạo multipart upload: %w", err)
// 	}

// 	// Xử lý lỗi và hủy tải lên nếu cần
// 	defer func() {
// 		if err != nil {
// 			client.AbortMultipartUpload(ctx, global.Config.MinIO.BUCKET_NAME, objectName, uploadID)
// 		}
// 	}()

// 	var parts []minio.CompletePart
// 	partNumber := 1
// 	chunkSize := int64(5 * 1024 * 1024) // 5MB chunks

// 	// Đọc và upload từng phần
// 	for {
// 		// Đọc chunk
// 		buf := make([]byte, chunkSize)
// 		n, err := src.Read(buf)
// 		if err != nil && err != io.EOF {
// 			return "", fmt.Errorf("lỗi đọc dữ liệu: %w", err)
// 		}

// 		if n == 0 {
// 			break
// 		}

// 		// Upload chunk
// 		part, err := client.PutObjectPart(ctx, global.Config.MinIO.BUCKET_NAME, objectName, uploadID, partNumber, bytes.NewReader(buf[:n]), int64(n), minio.PutObjectPartOptions{})
// 		if err != nil {
// 			return "", fmt.Errorf("lỗi khi tải lên phần %d: %w", partNumber, err)
// 		}

// 		parts = append(parts, minio.CompletePart{
// 			PartNumber: part.PartNumber,
// 			ETag:       part.ETag,
// 		})
// 		partNumber++

// 		if err == io.EOF {
// 			break
// 		}
// 	}

// 	// Hoàn tất quá trình multipart upload
// 	_, err = client.CompleteMultipartUpload(ctx, global.Config.MinIO.BUCKET_NAME, objectName, uploadID, parts)
// 	if err != nil {
// 		return "", fmt.Errorf("lỗi khi hoàn tất multipart upload: %w", err)
// 	}

// 	// Trả về URL của file đã tải lên
// 	protocol := "http"
// 	if global.Config.MinIO.USESSL {
// 		protocol = "https"
// 	}

// 	fileURL := fmt.Sprintf("%s://%s/%s/%s",
// 		protocol,
// 		global.Config.MinIO.ENDPOINT,
// 		global.Config.MinIO.BUCKET_NAME,
// 		objectName,
// 	)

// 	return fileURL, nil
// }
