package main

import (
	"Check/router"
)

func main() {
	h := router.InitHertzServer()
	router.InitRouter(h)
	h.Spin()

}
