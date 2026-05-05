package cache

import (
	"Order/model"
	"Order/utils"
	"context"
	"errors"
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

func ReduceStock(ctx context.Context, order model.OrderData) error {
	luaScript := `
	local stock = tonumber(redis.call('get',KEYS[1]) or 0)
	if stock > 0 then
		return redis.call('decr',KEYS[1])
	end
	return -1
	`
	res, err := redisDB.Eval(ctx, luaScript, []string{"stock:goods:" + strconv.Itoa(order.ProductId)}).Result()
	if err != nil {
		log.Println(err)
		return err
	}
	if res.(int64) == -1 {
		log.Println("库存不足！扣除失败！")
		return errors.New("库存不足！错误！")
	}
	log.Println("扣除成功！")
	return nil
}

//func WriteStock(ctx context.Context, productData model.ProductModel) error {
//	err := redisDB.Set(ctx, "stock:goods:"+strconv.Itoa(productData.Id), productData.Stock, 0).Err()
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	return nil
//}
//func WriteProduct(ctx context.Context, productData model.ProductModel) error {
//	productJSON, err := json.Marshal(productData)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	err1 := redisDB.Set(ctx, "info:goods:"+strconv.Itoa(productData.Id), productJSON, 0).Err()
//	if err1 != nil {
//		log.Println(err1)
//		return err1
//	}
//	return nil
//}
