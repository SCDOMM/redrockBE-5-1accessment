package api

import (
	"Order/model"
	"Order/sv"
	"Order/utils"
	"context"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
)

func OrderHandler(ctx context.Context, c *app.RequestContext) {
	orderData := model.OrderData{}
	err := c.BindJSON(&orderData)
	if err != nil {
		c.JSON(400, utils.FinalResponse{
			Status: "400",
			Info:   err.Error(),
			Data:   nil,
		})
		log.Println(err.Error())
	}
	err = sv.OrderHandler(ctx, orderData)
	if err != nil {
		c.JSON(400, utils.FinalResponse{
			Status: "400",
			Info:   "库存已经清空！",
			Data:   nil,
		})
		log.Println(err.Error())
	}
}
