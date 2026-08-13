package redis

import (
	"Order/utils"
	"context"
	"log"
	"strconv"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	redisDB *redis.Client
	once    sync.Once
)

func init() {
	once.Do(func() {
		config := utils.GetRedisConfig()
		rdb := redis.NewClient(&redis.Options{
			Addr:     config.Host + ":" + strconv.Itoa(config.Port),
			Password: config.Password,
			DB:       config.DB,
		})
		pong, err := rdb.Ping(context.Background()).Result()
		if err != nil {
			log.Println(err.Error())
			rdb.Close()
			return
		}
		log.Println("Redis连接成功！" + pong)
		redisDB = rdb
	})
}
