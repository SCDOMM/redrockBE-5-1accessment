package router

import (
	"Check/api"
	"GeneralConfig"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitHertzServer() *server.Hertz {
	config := GeneralConfig.GetHertzConfig()
	return server.Default(server.WithHostPorts(config.Host + ":" + strconv.Itoa(config.Port)))
}
func InitRouter(h *server.Hertz) {
	h.POST("/check", api.RouterHandler)
}
