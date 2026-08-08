package rd

import (
	"GeneralConfig"
	"Storage/model"
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
)

var redisDB *redis.Client

func InitRedis() error {
	config := GeneralConfig.GetRedisConfig()
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

func DecodeID(str string) string {
	return strings.Split(str, ":")[2]
}
func WriteStock(ctx context.Context, productData model.ProductModel) error {
	err := redisDB.Set(ctx, "stock:goods:"+strconv.Itoa(productData.Id), productData.Stock, 0).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func WriteProduct(ctx context.Context, productData model.ProductModel) error {
	productJSON, err := json.Marshal(productData)
	if err != nil {
		log.Println(err)
		return err
	}
	err1 := redisDB.Set(ctx, "info:goods:"+strconv.Itoa(productData.Id), productJSON, 0).Err()
	if err1 != nil {
		log.Println(err1)
		return err1
	}
	return nil
}
