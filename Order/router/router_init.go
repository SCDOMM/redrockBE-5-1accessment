package router

import (
	"GeneralConfig"
	"Order/api"
	"strconv"
	"sync"

	"github.com/cloudwego/hertz/pkg/app/server"
)

var (
	once sync.Once
	h    *server.Hertz
)

func init() {
	once.Do(func() {
		config := GeneralConfig.GetHertzConfig()
		h = server.Default(server.WithHostPorts(config.Host + ":" + strconv.Itoa(config.Port)))
	})
}

func InitRouter() {
	h.POST("/order", api.OrderHandler)
}
