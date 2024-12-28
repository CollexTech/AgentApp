package datastore

import (
	"context"
	"fmt"

	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var RateLimiter *redis_rate.Limiter

func InitializeRateLimiter() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", viper.GetString("REDIS_HOST"), viper.GetString("REDIS_PORT")),
		DB:   1,
	})
	_ = rdb.FlushDB(ctx).Err()

	RateLimiter = redis_rate.NewLimiter(rdb)
}
