package middlewares

import (
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

// RateLimiterMiddlewareRedis giới hạn tốc độ request bằng Redis
func RateLimiterPrivateMiddlewareRedis() gin.HandlerFunc {
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
		limit := consts.RATELIMIT_REQUEST_PRIVATE
		expiration := time.Duration(consts.RATELIMIT_SECOND__PRIVATE) * time.Second

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

func RateLimiterGlobalMiddlewareRedis() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		client := global.Rdb
		key := "ratelimit:global"

		// Sử dụng Redis Pipeline để tối ưu hiệu suất
		pipe := client.TxPipeline()

		// Tăng số request hiện tại
		count := pipe.Incr(ctx, key)
		pipe.Expire(ctx, key, time.Duration(consts.RATELIMIT_SECOND__GLOBAL)*time.Second)

		_, err := pipe.Exec(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis error"})
			c.Abort()
			return
		}

		// Nếu request vượt quá giới hạn, trả về lỗi 429
		if count.Val() > int64(consts.RATELIMIT_REQUEST_GLOBAL) {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
			c.Abort()
			return
		}

		c.Next()
	}
}
