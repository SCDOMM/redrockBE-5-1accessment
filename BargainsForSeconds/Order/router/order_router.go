package router

import (
	"Order/api"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitRouter(h *server.Hertz) {
	h.POST("/order", api.OrderHandler)
}
