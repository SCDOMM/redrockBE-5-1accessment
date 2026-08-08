package api

import (
	"Check/model"
	"Check/sv"
	"Check/utils"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

func RouterHandler(ctx context.Context, c *app.RequestContext) {
	orderData := model.OrderData{}
	err := c.BindJSON(&orderData)
	if err != nil {
		c.JSON(400, utils.FinalResponse{
			Status: "400",
			Info:   err.Error(),
			Data:   nil,
		})
	}
	err = sv.CheckHandler(orderData)
	if err != nil {
		c.JSON(400, utils.FinalResponse{
			Status: "400",
			Info:   err.Error(),
			Data:   nil,
		})
	}
	c.JSON(200, utils.FinalResponse{
		Status: "200",
		Info:   "Order Success!",
	})
}
