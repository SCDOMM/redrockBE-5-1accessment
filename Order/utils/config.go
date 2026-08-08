package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"gopkg.in/yaml.v2"
)

var (
	etcdConfig *EtcdConfig
	center     *Center
	configOnce sync.Once
)

// EtcdConfig 用于连接 etcd 的配置
type EtcdConfig struct {
	Key  string `yaml:"key"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// 业务配置结构体

type RabbitMQConfig struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Vhost    string `yaml:"vhost"`
}
type MySQLConfig struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
}
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}
type MachineId struct {
	Id int64 `yaml:"id"`
}
type HertzConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
type KitexConfig struct {
	ServerName string `yaml:"server_name"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
}
type ConfigData struct {
	RabbitMQConfig RabbitMQConfig `yaml:"RabbitMQConfig"`
	MySQLConfig    MySQLConfig    `yaml:"MySQLConfig"`
	RedisConfig    RedisConfig    `yaml:"RedisConfig"`
	MachineId      MachineId      `yaml:"MachineId"`
	HertzConfig    HertzConfig    `yaml:"HertzConfig"`
	KitexConfig    KitexConfig    `yaml:"KitexConfig"`
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

		endpoints := []string{fmt.Sprintf("%s:%d", etcdConfig.Host, etcdConfig.Port)}
		if err := initCenter(endpoints, etcdConfig.Key); err != nil {
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
		etcdConfig = &EtcdConfig{Key: key, Host: host, Port: port}
		return nil
	}

	// 拉本地配置文件
	configPath := os.Getenv("CONFIG_FILE")
	if configPath == "" {
		configPath = "./config.yaml"
	}
	dataBytes, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("读取本地配置文件失败: %v", err)
	}

	etcdConfig = &EtcdConfig{}
	if err := yaml.Unmarshal(dataBytes, etcdConfig); err != nil {
		return fmt.Errorf("解析本地配置文件失败: %v", err)
	}

	if etcdConfig.Key == "" || etcdConfig.Host == "" || etcdConfig.Port == 0 {
		return fmt.Errorf("本地配置文件中 EtcdConfig 不完整")
	}
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

// 对外暴露的配置读取函数

// GetEtcdConfig 返回 etcd 连接配置
func GetEtcdConfig() *EtcdConfig {
	return etcdConfig
}

func GetMySQLConfig() MySQLConfig {
	center.mu.RLock()
	defer center.mu.RUnlock()
	return center.config.MySQLConfig
}

func GetRedisConfig() RedisConfig {
	center.mu.RLock()
	defer center.mu.RUnlock()
	return center.config.RedisConfig
}

func GetKitexConfig() KitexConfig {
	center.mu.RLock()
	defer center.mu.RUnlock()
	return center.config.KitexConfig
}

func GetHertzConfig() HertzConfig {
	center.mu.RLock()
	defer center.mu.RUnlock()
	return center.config.HertzConfig
}

func GetRabbitMQConfig() RabbitMQConfig {
	center.mu.RLock()
	defer center.mu.RUnlock()
	return center.config.RabbitMQConfig
}

func GetMachineId() MachineId {
	center.mu.RLock()
	defer center.mu.RUnlock()
	return center.config.MachineId
}
