package router

import (
	"Check/api"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitRouter(h *server.Hertz) {
	h.POST("/check", api.RouterHandler)
}
