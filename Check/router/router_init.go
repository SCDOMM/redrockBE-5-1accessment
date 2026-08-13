package router

import (
	"Check/api"
	"Check/utils"
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
		config := utils.GetHertzConfig()
		h = server.Default(server.WithHostPorts(config.Host + ":" + strconv.Itoa(config.Port)))
	})
}
func InitRouter() {
	h.POST("/check", api.RouterHandler)
	h.Spin()
}
