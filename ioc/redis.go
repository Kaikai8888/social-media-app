package ioc

import (
	"social-media-app/config"

	"github.com/redis/go-redis/v9"
)

func InitRedis() redis.Cmdable {
	return redis.NewClient(&redis.Options{
		Addr: config.Config.RedisConfig.Addr,
	})
}
