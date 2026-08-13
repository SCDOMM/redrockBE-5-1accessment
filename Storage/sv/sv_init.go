package sv

import (
	"Storage/dao"
	"Storage/sv/rd"
	"context"
	"log"
)

// 初始化商品
func init() {
	products, err := dao.CheckProduct()
	if err != nil {
		log.Println(err)
		return
	}
	ctx := context.Background()

	for _, product := range *products {
		err = rd.WriteProduct(ctx, product)
		if err != nil {
			log.Println(err)
			continue
		}
		err = rd.WriteStock(ctx, product)
		if err != nil {
			log.Println(err)
		}
	}
}
