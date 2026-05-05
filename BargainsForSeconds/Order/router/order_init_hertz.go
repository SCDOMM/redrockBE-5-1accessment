package router

import (
	"Order/utils"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitHertz() *server.Hertz {
	config := utils.GetHertzConfig()
	return server.Default(server.WithHostPorts(config.Host + ":" + strconv.Itoa(config.Port)))
}
