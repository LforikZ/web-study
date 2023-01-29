package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"web-study/settings"
)

var (
	client *redis.Client
)

func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize, //最大连接数
	})
	_, err = client.Ping().Result()
	return
}

func Close() {
	client.Close()
}
