package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func getOptions(host string, port int, db int, password string) *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       db,
	}
}

// NewRedis 创建 Redis 客户端连接（保留供兼容，推荐使用 driver 包）
func NewRedis(host string, port int, db int, password string) (*redis.Client, error) {
	client := redis.NewClient(getOptions(host, port, db, password))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
