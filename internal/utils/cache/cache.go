package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"go-backend-api/global"
	"log"

	"github.com/redis/go-redis/v9"
)

func GetCache(ctx context.Context, key string, obj interface{}) error {
	rs, err := global.Rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return fmt.Errorf("key %s not found", key)
	} else if err != nil {
		return err
	}
	log.Println("🔍 Dữ liệu từ Redis:", rs)
	if err := json.Unmarshal([]byte(rs), obj); err != nil {
		return fmt.Errorf("failed to unmarshal")
	}
	return nil
}
