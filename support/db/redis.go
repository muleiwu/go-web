package db

import (
	"fmt"
	"sync"

	"github.com/redis/go-redis/v9"
	"mliev.com/template/go-web/support"
)

var (
	rdb  *redis.Client
	once sync.Once
)

func GetRedis() *redis.Client {
	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d",
				support.Env("redis.host", "localhost").(string),
				support.Env("redis.port", 6379).(int),
			),
			Password: support.Env("redis.password", "").(string), // no password set
			DB:       support.Env("redis.db", 0).(int),           // use default DB
		})
	})

	return rdb
}
