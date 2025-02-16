package util

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func ClearRedis(t *testing.T, redis redis.Cmdable) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	if err := redis.FlushDB(ctx).Err(); err != nil {
		assert.FailNow(t, "failed to clear redis", err.Error())
	}
}
