package utils

type LocalConfig struct {
	EtcdConfig     EtcdConfig     `yaml:"EtcdConfig"`
	ResolverConfig ResolverConfig `yaml:"ResolverConfig"`
	RegisterConfig RegisterConfig `yaml:"RegisterConfig"`
}
type RegistryConfig struct {
	RegisterConfig RegisterConfig `yaml:"RegisterConfig"`
	ResolverConfig ResolverConfig `yaml:"ResolverConfig"`
}
type ResolverConfig struct {
	NameSpace   string `yaml:"namespace"`
	Env         string `yaml:"env"`
	ServiceName string `yaml:"service_name"`
}
type RegisterConfig struct {
	NameSpace   string `yaml:"namespace"`
	Env         string `yaml:"env"`
	ServiceName string `yaml:"service_name"`
	InstanceID  string `yaml:"instance_id"`
	Address     string `yaml:"address"`
	Ttl         int    `yaml:"ttl"`
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

// 对外暴露的配置读取函数

// GetEtcdConfig 返回 etcd 连接配置
func GetEtcdConfig() *EtcdConfig {
	return EtcdConfigSample
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
func GetRegisterConfig() RegisterConfig {
	return RegistryConfigSample.RegisterConfig
}
func GetResolverConfig() ResolverConfig {
	return RegistryConfigSample.ResolverConfig
}
