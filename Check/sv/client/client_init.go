package client

import (
	"Check/kitex_gen/checkserver/service/checkservice"
	"Check/registry"
	"Check/utils"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cloudwego/kitex/client"
)

var (
	kitexClient atomic.Value
	mu          sync.RWMutex
	lastAddr    string
)

func init() {
	// 初始化服务发现
	registry.InitDiscovery()
	// 等待 A 服务注册
	addr := waitForAddr()
	// 创建 Kitex Client
	updateClient(addr)
	// 后台周期性检查地址变化并更新 Client
	go func() {
		for {
			time.Sleep(5 * time.Second)
			if addr := registry.GetCurrentAddr(); addr != "" {
				updateClient(addr)
			}
		}
	}()
}
func getKitexClient() checkservice.Client {
	v := kitexClient.Load()
	if v == nil {
		return nil
	}
	return v.(checkservice.Client)
}
func updateClient(addr string) {
	mu.Lock()
	defer mu.Unlock()
	if addr == lastAddr && getKitexClient() != nil {
		return
	}
	lastAddr = addr
	cli := checkservice.MustNewClient(
		utils.GetResolverConfig().ServiceName,
		client.WithHostPorts(addr),
	)
	kitexClient.Store(cli)
	log.Printf("Kitex Client 已连接: %s", addr)
}
func waitForAddr() string {
	for {
		if addr := registry.GetCurrentAddr(); addr != "" {
			return addr
		}
		log.Println("等待 Storage 服务注册...")
		time.Sleep(2 * time.Second)
	}
}
