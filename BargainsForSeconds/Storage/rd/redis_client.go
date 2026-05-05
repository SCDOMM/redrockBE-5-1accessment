package rd

import (
	"Storage/utils"
	"context"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var redisDB *redis.Client

func InitRedis() error {
	config := utils.GetRedisConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + strconv.Itoa(config.Port),
		Password: config.Password,
		DB:       config.DB,
	})
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println("连接成功！" + pong)
	redisDB = rdb
	return nil
}
