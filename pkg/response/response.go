package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`    //status code
	Message string      `json:"message"` // Thông báo lỗi
	Data    interface{} `json:"data"`    //dữ liệu được return
}

type ErrorResponseData struct {
	Code   int         `json:"code"`   //status code
	Err    string      `json:"error"`  // Thông báo lỗi
	Detail interface{} `json:"detail"` //dữ liệu được return
}

// success response
func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

// Error response
func ErrorResponse(c *gin.Context, code int, message string) {
	if message == "" {
		message = msg[code]
	}
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
