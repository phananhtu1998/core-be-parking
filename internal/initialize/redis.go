package initialize

import (
	"context"
	"fmt"
	"go-backend-api/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.Database, // use default DB
		PoolSize: 10,         // có 10 connect cho mỗi cpu khả dụng
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization Error: %v", zap.Error(err))
	}
	fmt.Println("InitRedis is running")
	global.Rdb = rdb
	//redisExample()
}

// func redisExample() {
// 	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
// 	if err != nil {
// 		fmt.Println("Error redis setting:", zap.Error(err))
// 		return
// 	}
// 	value, err := global.Rdb.Get(ctx, "score").Result()
// 	if err != nil {
// 		fmt.Println("Error redis setting:", zap.Error(err))
// 		return
// 	}
// 	global.Logger.Info("Value score is: ", zap.String("score", value))
// }
