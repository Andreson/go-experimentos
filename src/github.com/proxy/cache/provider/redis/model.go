package redis

import (
	"github.com/go-redis/redis/v8"
)

type RedisOperations struct {
	ConneError error
	Client     *redis.Client
}

type RedisModel struct {
	ResponseCmd   *redis.StringCmd
	IntCmd        *redis.IntCmd
	StatusCmd     *redis.StatusCmd
	Success       bool
	Err           error
	RunOperations bool
}

type RedisProvider struct {
	ConneError error
	Client     *redis.Client
}
