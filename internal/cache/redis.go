package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"order_machine/configs"
)

func RedisConnFactory(redisConfig *configs.RedisConf, db int) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       db,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}
