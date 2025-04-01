package upload

import (
	"fmt"
	"go-backend-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadFile
// @Summary Upload file
// @Description API upload file cho hệ thống
// @Tags Upload
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Security     ApiKeyAuth
// @Param file formData file true "File cần upload"
// @Success 200 {object} response.ResponseData
// @Failure 500 {object} response.ErrorResponseData "Server error"
// @Router /upload/upload_file [post]
func UploadFileHandler(c *gin.Context) {
	// Lấy file từ request
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Không thể lấy file từ request"})
		return
	}

	// Mở file
	fileContent, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể mở file"})
		return
	}
	defer fileContent.Close()
	fileURL, err := service.UploadItem().UploadFile(fileContent, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Lỗi upload file: %v", err)})
		return
	}
	// Trả về kết quả
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Upload thành công",
		"url":     fileURL,
	})
}
