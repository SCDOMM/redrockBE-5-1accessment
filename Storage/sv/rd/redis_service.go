package rd

import (
	"Storage/model"
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

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
