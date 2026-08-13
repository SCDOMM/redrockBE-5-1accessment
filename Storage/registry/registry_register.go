package registry

import (
	"Storage/utils"
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	registerOnce sync.Once
	registration *Registration
)

type Registry struct {
	cli *clientv3.Client
}

// Registration 单个服务实例的注册句柄
type Registration struct {
	cli      *clientv3.Client
	leaseID  clientv3.LeaseID
	key      string
	stop     chan struct{}
	stopOnce sync.Once
}

func InitRegister() {
	registerOnce.Do(func() {
		reg, err := NewRegistry(utils.GetEtcdEndpoints())
		if err != nil {
			log.Fatalf("[registry] 连接 etcd 失败: %v", err)
		}
		cfg := utils.GetRegisterConfig()
		if cfg.Address == "" {
			cfg.Address = fmt.Sprintf("%s:%d", utils.GetKitexConfig().Host, utils.GetKitexConfig().Port)
		}
		if cfg.InstanceID == "" {
			cfg.InstanceID = fmt.Sprintf("%s-%d", cfg.ServiceName, os.Getpid())
		}
		registration, err = reg.Register(
			cfg.NameSpace,
			cfg.Env,
			cfg.ServiceName,
			cfg.InstanceID,
			cfg.Address,
			int64(cfg.Ttl),
		)
		if err != nil {
			log.Fatalf("[registry] 服务注册失败: %v", err)
		}
	})
}

func NewRegistry(endpoints []string) (*Registry, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return &Registry{cli: cli}, nil
}

// Client 暴露底层 etcd client，供 Discovery 复用同一个连接
func (r *Registry) Client() *clientv3.Client { return r.cli }

// Stop 注销服务（撤销租约，etcd 自动删除节点）
func (r *Registration) Stop() {
	r.stopOnce.Do(func() { close(r.stop) })
}
func CloseRegistry() {
	if registration != nil {
		registration.Stop()
	}
}

// Register 注册服务实例，后台自动心跳续约
func (r *Registry) Register(namespace, env, serviceName, instanceID, addr string, ttl int64) (*Registration, error) {
	ctx := context.Background()

	// TTL秒后不续约则自动过期
	grant, err := r.cli.Grant(ctx, ttl)
	if err != nil {
		return nil, fmt.Errorf("创建租约失败: %w", err)
	}

	// 写入节点/myapp/dev/services/service-a/实例ID
	key := fmt.Sprintf("/%s/%s/services/%s/%s", namespace, env, serviceName, instanceID)
	if _, err := r.cli.Put(ctx, key, addr, clientv3.WithLease(grant.ID)); err != nil {
		r.cli.Revoke(ctx, grant.ID)
		return nil, fmt.Errorf("写入节点失败: %w", err)
	}

	// 心跳续约
	keepAliveCh, err := r.cli.KeepAlive(ctx, grant.ID)
	if err != nil {
		r.cli.Revoke(ctx, grant.ID)
		return nil, fmt.Errorf("启动续约失败: %w", err)
	}

	reg := &Registration{
		cli:     r.cli,
		leaseID: grant.ID,
		key:     key,
		stop:    make(chan struct{}),
	}
	go func() {
		for {
			select {
			case _, ok := <-keepAliveCh:
				if !ok { // 续约通道异常关闭
					log.Printf("[registry] 服务 %s 心跳续约中断", key)
					return
				}
			case <-reg.stop: // 主动注销
				r.cli.Revoke(context.Background(), grant.ID)
				log.Printf("[registry] 服务已注销: %s", key)
				return
			}
		}
	}()

	log.Printf("[registry] 服务注册成功: %s -> %s (TTL=%ds)", key, addr, ttl)
	return reg, nil
}
