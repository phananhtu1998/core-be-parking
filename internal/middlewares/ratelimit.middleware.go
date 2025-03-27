package middlewares

import (
	"context"
	"fmt"
	"go-backend-api/global"
	consts "go-backend-api/internal/const"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

// RateLimiterMiddlewareRedis giới hạn tốc độ request bằng Redis
func RateLimiterMiddlewareRedis() gin.HandlerFunc {
	return func(c *gin.Context) {
		client := global.Rdb
		endpoint := c.FullPath() // Lấy đường dẫn API
		ip := c.ClientIP()       // Lấy địa chỉ IP của user
		key := fmt.Sprintf("ratelimit:%s:%s", endpoint, ip)

		// Lấy số lượng request hiện tại
		count, _ := client.Get(ctx, key).Int()

		// Giới hạn request (ví dụ: 5 request mỗi 10 giây)
		limit := consts.RATELIMIT_REQUEST
		expiration := time.Duration(consts.RATELIMIT_SECOND) * time.Second

		if count >= limit {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
			c.Abort()
			return
		}
		// Tăng số request và đặt thời gian hết hạn Expire key sau 10 giây
		client.Incr(ctx, key)
		client.Expire(ctx, key, expiration)
		c.Next()
	}
}
