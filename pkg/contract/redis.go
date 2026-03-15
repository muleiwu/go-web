package contract

import "github.com/redis/go-redis/v9"

// RedisClient Redis 服务契约（后续可抽象为接口）
type RedisClient = redis.Client
