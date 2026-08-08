package GeneralConfig
import (
	"fmt"
	"os"
	"sync"
	"gopkg.in/yaml.v3"
)

var (
	configData ConfigData
	once       sync.Once
)
var _ = yaml.Unmarshal

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

func init() {
	once.Do(func() {
		dataBytes, err := os.ReadFile("../GeneralConfig/config.yaml")
		if err != nil {
			fmt.Println("读取配置失败！" + err.Error())
		}
		err = yaml.Unmarshal(dataBytes, &configData)
		if err != nil {
			fmt.Println("解析配置失败！" + err.Error())
		}
	})
}
func GetRabbitMQConfig() RabbitMQConfig {
	return configData.RabbitMQConfig
}
func GetMySQLConfig() MySQLConfig {
	return configData.MySQLConfig
}
func GetRedisConfig() RedisConfig {
	return configData.RedisConfig
}
func GetMachineId() int64 {
	return configData.MachineId.Id
}
func GetHertzConfig() HertzConfig {
	return configData.HertzConfig
}
func GetKitexConfig() KitexConfig {
	return configData.KitexConfig
}
