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
		return
	}
	err = sv.OrderHandler(ctx, orderData)
	if err != nil {
		c.JSON(500, utils.FinalResponse{
			Status: "500",
			Info:   "库存已经清空！",
			Data:   nil,
		})
		log.Println(err.Error())
		return
	}
	c.JSON(200, utils.FinalResponse{
		Status: "200",
		Info:   "下单成功！",
		Data:   nil,
	})
}
