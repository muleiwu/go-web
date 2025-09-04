package impl

import (
	"context"
	"fmt"
	"time"

	"cnb.cool/mliev/examples/go-web/internal/interfaces"
	"github.com/redis/go-redis/v9"
)

func getOptions(host string, port int, db int, password string) *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       db,
	}
}

func NewRedis(Helper interfaces.HelperInterface, host string, port int, db int, password string) (*redis.Client, error) {

	// 创建Redis客户端
	client := redis.NewClient(getOptions(host, port, db, password))

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
