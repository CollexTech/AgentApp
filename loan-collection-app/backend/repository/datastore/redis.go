package datastore

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func InitRedisClient() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", viper.GetString("REDIS_HOST"), viper.GetString("REDIS_PORT")),
	})
}
