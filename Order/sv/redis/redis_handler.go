package redis

import (
	"Order/model"
	"context"
	"fmt"
	"log"
	"strconv"
)

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
		return fmt.Errorf("库存不足，扣除失败")
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
