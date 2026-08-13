package registry

import (
	"Check/utils"
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Discovery struct {
	cli         *clientv3.Client
	namespace   string
	env         string
	serviceName string
	mu          sync.RWMutex
	nodes       map[string]string // instanceID -> addr
}

var (
	discovery    *Discovery
	currentAddr  atomic.Value
	resolverOnce sync.Once
	rrIndex      uint32
)

func InitDiscovery() {
	resolverOnce.Do(func() {
		cli, err := clientv3.New(clientv3.Config{
			Endpoints:   utils.GetEtcdEndpoints(),
			DialTimeout: 5 * time.Second,
		})
		if err != nil {
			log.Fatalf("[discovery] 连接 etcd 失败: %v", err)
		}
		cfg := utils.GetResolverConfig()
		d, err := NewDiscovery(cli, cfg.NameSpace, cfg.Env, cfg.ServiceName)
		if err != nil {
			log.Fatalf("[discovery] 初始化失败: %v", err)
		}
		discovery = d
		// 先选一个可用地址
		if addr, err := d.Pick(); err == nil {
			currentAddr.Store(addr)
		}
		// 后台 watch，自动更新 currentAddr
		go d.watchAndUpdate()
	})
}

func NewDiscovery(cli *clientv3.Client, namespace, env, serviceName string) (*Discovery, error) {
	d := &Discovery{
		cli:         cli,
		namespace:   namespace,
		env:         env,
		serviceName: serviceName,
		nodes:       make(map[string]string),
	}
	if err := d.initNodes(); err != nil { // 首次全量拉取
		return nil, err
	}
	return d, nil
}

func (d *Discovery) prefix() string {
	return fmt.Sprintf("/%s/%s/services/%s/", d.namespace, d.env, d.serviceName)
}

func (d *Discovery) initNodes() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := d.cli.Get(ctx, d.prefix(), clientv3.WithPrefix())
	if err != nil {
		return err
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	for _, kv := range resp.Kvs {
		id := strings.TrimPrefix(string(kv.Key), d.prefix())
		d.nodes[id] = string(kv.Value)
	}
	return nil
}

func (d *Discovery) watchAndUpdate() {
	rch := d.cli.Watch(context.Background(), d.prefix(), clientv3.WithPrefix())
	log.Printf("[registry] 开始监听服务 %s 节点变化...", d.serviceName)
	for wResp := range rch {
		for _, ev := range wResp.Events {
			id := strings.TrimPrefix(string(ev.Kv.Key), d.prefix())
			d.mu.Lock()
			switch ev.Type {
			case clientv3.EventTypePut:
				d.nodes[id] = string(ev.Kv.Value)
				log.Printf("[registry] 节点上线: %s -> %s", ev.Kv.Key, ev.Kv.Value)
			case clientv3.EventTypeDelete:
				delete(d.nodes, id)
				log.Printf("[registry] 节点下线: %s", ev.Kv.Key)
			}
			d.mu.Unlock()
		}
		// 节点变化后重新 Pick 一次
		if addr, err := d.Pick(); err == nil {
			currentAddr.Store(addr)
			log.Printf("[discovery] 当前可用节点更新: %s", addr)
		} else {
			currentAddr.Store("")
		}
	}
}

// Resolve 返回当前所有可用节点地址
func (d *Discovery) Resolve() []string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	list := make([]string, 0, len(d.nodes))
	for _, addr := range d.nodes {
		list = append(list, addr)
	}
	return list
}

// Pick 简单轮询选一个节点（A 多副本时天然负载均衡）
func (d *Discovery) Pick() (string, error) {
	nodes := d.Resolve()
	if len(nodes) == 0 {
		return "", fmt.Errorf("服务 %s 当前没有可用节点", d.serviceName)
	}
	i := atomic.AddUint32(&rrIndex, 1)
	return nodes[int(i)%len(nodes)], nil
}

func GetCurrentAddr() string {
	addr, _ := currentAddr.Load().(string)
	return addr
}
