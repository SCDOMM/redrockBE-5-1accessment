package router

import (
	"GeneralConfig"
	"Order/api"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitHertzServer() *server.Hertz {
	config := GeneralConfig.GetHertzConfig()
	return server.Default(server.WithHostPorts(config.Host + ":" + strconv.Itoa(config.Port)))
}
func InitRouter(h *server.Hertz) {
	h.POST("/order", api.OrderHandler)
}
