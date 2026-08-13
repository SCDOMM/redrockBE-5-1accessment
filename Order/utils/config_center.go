package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"gopkg.in/yaml.v3"
)

var (
	EtcdConfigSample *EtcdConfig
	center           *Center
	configOnce sync.Once
)

// EtcdConfig 用于连接 etcd 的配置
type EtcdConfig struct {
	Key  string `yaml:"key"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// Center 配置中心客户端
type Center struct {
	cli    *clientv3.Client
	key    string
	mu     sync.RWMutex
	config ConfigData
}

func init() {
	configOnce.Do(func() {
		if err := loadEtcdConfig(); err != nil {
			log.Fatalf("加载 etcd 配置失败: %v", err)
		}

		endpoints := []string{fmt.Sprintf("%s:%d", EtcdConfigSample.Host, EtcdConfigSample.Port)}
		if err := initCenter(endpoints, EtcdConfigSample.Key); err != nil {
			log.Fatalf("初始化配置中心失败: %v", err)
		}
		go center.Watch(context.Background())
	})
}

// loadEtcdConfig 从环境变量或本地文件获取 etcd 连接信息
func loadEtcdConfig() error {
	//拉环境变量
	host := os.Getenv("ETCD_HOST")
	portStr := os.Getenv("ETCD_PORT")
	key := os.Getenv("ETCD_KEY")

	if host != "" && portStr != "" && key != "" {
		port := 0
		if _, err := fmt.Sscanf(portStr, "%d", &port); err != nil {
			return fmt.Errorf("解析 ETCD_PORT 失败: %v", err)
		}
		EtcdConfigSample = &EtcdConfig{Key: key, Host: host, Port: port}
		return nil
	}

	localConfig := LocalConfig{}
	dataBytes, err := os.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println("读取配置失败！" + err.Error())
	}
	err = yaml.Unmarshal(dataBytes, &localConfig)
	if err != nil {
		fmt.Println("解析配置失败！" + err.Error())
	}
	EtcdConfigSample = &localConfig.EtcdConfig
	return nil
}

// initCenter 连接 etcd、拉取配置、启动监听
func initCenter(endpoints []string, key string) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return err
	}

	c := &Center{cli: cli, key: key}
	if err := c.load(); err != nil {
		cli.Close()
		return err
	}

	center = c

	log.Printf("配置中心初始化完成，key: %s，endpoints: %v", key, endpoints)
	return nil
}

// load 从 etcd 拉取并解析配置
func (c *Center) load() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := c.cli.Get(ctx, c.key)
	if err != nil {
		return err
	}
	if len(resp.Kvs) == 0 {
		return fmt.Errorf("key %q 不存在", c.key)
	}

	var data ConfigData
	if err := yaml.Unmarshal(resp.Kvs[0].Value, &data); err != nil {
		return err
	}

	c.mu.Lock()
	c.config = data
	c.mu.Unlock()

	log.Printf("配置已加载: %s\n", c.key)
	return nil
}

// Watch 监听配置变更，自动热更新
func (c *Center) Watch(ctx context.Context) {
	rch := c.cli.Watch(ctx, c.key)
	for wResp := range rch {
		for _, ev := range wResp.Events {
			if ev.Type == clientv3.EventTypePut {
				if err := c.load(); err != nil {
					log.Printf("热更新失败: %v", err)
				} else {
					log.Println("配置已热更新")
				}
			}
		}
	}
}
