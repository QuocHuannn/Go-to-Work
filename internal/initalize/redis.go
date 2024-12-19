package initalize

import (
	"context"
	"fmt"
	"github.com/QuocHuannn/Go-to-Work/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.Database, // use default DB
		PoolSize: 10,         // connection pool size
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect redis", zap.Error(err))
	}
	fmt.Println("Init redis is running")
	global.Rdb = rdb
}

func redisExample() {
	err := global.Rdb.Set(ctx, "Score", 100, 0).Err()
	if err != nil {
		fmt.Println("Failed to set key Score", zap.Error(err))
		return
	}
	value, err := global.Rdb.Get(ctx, "Score").Result()
	if err != nil {
		fmt.Println("Failed to get key Score", zap.Error(err))
		return
	}
	global.Logger.Info("Value core is: ", zap.String("score", value))
}
