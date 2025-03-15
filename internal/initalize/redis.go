package initalize

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/QuocHuannn/Go-to-Work/global"
	"github.com/QuocHuannn/Go-to-Work/internal/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	// Use new config structure
	r := config.Cfg.Redis

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
		DB:       r.DB,       // Updated field name from Database to DB
		PoolSize: 10,         // connection pool size
	})

	fmt.Printf("Redis connection: %s:%d\n", host, port)

	// Thêm cơ chế retry cho kết nối Redis trong Docker
	maxRetries := 5
	retryDelay := 5 * time.Second
	var err error

	for i := 0; i < maxRetries; i++ {
		_, err = rdb.Ping(ctx).Result()
		if err == nil {
			global.Logger.Info("Redis connection successful", zap.String("host", host), zap.Int("port", port))
			break
		}

		fmt.Printf("Failed to connect to Redis (attempt %d/%d): %v\n", i+1, maxRetries, err)
		global.Logger.Warn("Failed to connect to Redis", zap.Error(err), zap.Int("attempt", i+1))

		if i < maxRetries-1 {
			fmt.Printf("Retrying in %v...\n", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	if err != nil {
		global.Logger.Error("Failed to connect redis after multiple attempts", zap.Error(err))
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
