package initialize

import (
	"context"
	"fmt"
	"go-cover-parroto/global"

	"github.com/redis/go-redis/v9"
)

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.DB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(fmt.Errorf("failed to connect to Redis: %w", err))
	}

	global.Rdb = rdb
	global.Logger.Info("Redis connected")
}