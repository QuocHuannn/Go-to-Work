package initalize

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/QuocHuannn/Go-to-Work/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis

	// Override host with environment variable if running in Docker
	host := r.Host
	if os.Getenv("REDIS_HOST") != "" {
		host = os.Getenv("REDIS_HOST")
		fmt.Printf("Using Redis host from environment: %s\n", host)
	} else {
		fmt.Printf("Using Redis host from config: %s\n", host)
	}

	// Get port from environment if available
	port := r.Port
	if os.Getenv("REDIS_PORT") != "" {
		if p, err := strconv.Atoi(os.Getenv("REDIS_PORT")); err == nil {
			port = p
			fmt.Printf("Using Redis port from environment: %d\n", port)
		}
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: r.Password, // no password set
		DB:       r.Database, // use default DB
		PoolSize: 10,         // connection pool size
	})

	fmt.Printf("Redis connection: %s:%d\n", host, port)

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect redis", zap.Error(err))
	} else {
		global.Logger.Info("Redis connection successful", zap.String("host", host), zap.Int("port", port))
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
