package middlewares

import (
	"context"
	"fmt"
	"go-backend-api/global"
	consts "go-backend-api/internal/const"
	"go-backend-api/internal/model"
	"go-backend-api/internal/utils/auth"
	"go-backend-api/internal/utils/cache"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

// RateLimiterMiddlewareRedis giới hạn tốc độ request bằng Redis
func RateLimiterMiddlewareRedis() gin.HandlerFunc {
	return func(c *gin.Context) {
		client := global.Rdb
		ctx := c.Request.Context()
		// Lấy token từ request
		jwtToken, _ := auth.ExtracBearerToken(c)
		claims, _ := auth.VerifyTokenSubject(jwtToken)
		var infoUser model.GetCacheToken
		if err := cache.GetCache(ctx, claims.Subject, &infoUser); err != nil {
			return
		}
		endpoint := c.FullPath() // Lấy đường dẫn API
		ip := c.ClientIP()       // Lấy địa chỉ IP của user
		// truyền cả Id vì giả sử trong 1 cty thì IP mạng có thể giống nhau nên kèm theo ID để phân biệt
		key := fmt.Sprintf("ratelimit:%s:%s:%", endpoint, ip, infoUser.ID)

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
