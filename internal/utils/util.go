package utils

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserKey(hashKey string) string {
	return fmt.Sprintf("u:%s:otp", hashKey)
}

func GenerateCliTokenUUID(userId int) string {
	newUUID := uuid.New()
	uuidString := strings.ReplaceAll((newUUID).String(), "", "")
	return strconv.Itoa(userId) + "clitoken" + uuidString
}

func CompareNullString(a, b sql.NullString) bool {
	if !a.Valid && !b.Valid {
		return true
	}
	if a.Valid != b.Valid {
		return false
	}
	return a.String == b.String
}

func ValidateJSONFields(ctx *gin.Context, inputStruct interface{}, validFields map[string]bool) error {
	// Đọc toàn bộ body request
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return errors.New("invalid request body")
	}

	// Chuyển body thành map để kiểm tra field dư
	var rawData map[string]interface{}
	if err := json.Unmarshal(body, &rawData); err != nil {
		return errors.New("invalid JSON format")
	}

	// Bind dữ liệu vào struct để kiểm tra định dạng
	if err := json.Unmarshal(body, inputStruct); err != nil {
		return errors.New("invalid JSON fields")
	}

	// Kiểm tra nếu có field dư
	for key := range rawData {
		if !validFields[key] {
			return errors.New("unexpected field: " + key)
		}
	}

	// Reset lại body request để các middleware khác có thể sử dụng
	ctx.Request.Body = io.NopCloser(io.MultiReader(io.NewSectionReader(bytes.NewReader(body), 0, int64(len(body)))))

	return nil
}
