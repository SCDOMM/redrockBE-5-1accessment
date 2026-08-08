package client

import (
	"Check/utils"
	"checkserver/kitex_gen/checkserver/service/checkservice"
	"strconv"

	"github.com/cloudwego/kitex/client"
)

func init() {
	config := utils.GetKitexConfig()
	kitexClient = checkservice.MustNewClient(config.ServerName, client.WithHostPorts(config.Host+strconv.Itoa(config.Port)))
}
