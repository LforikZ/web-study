package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"web-study/settings"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize, //最大连接数
	})
	_, err = rdb.Ping().Result()
	return
}

func Close() {
	rdb.Close()
}
