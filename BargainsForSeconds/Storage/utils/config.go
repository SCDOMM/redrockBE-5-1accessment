package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var configData ConfigData

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
type ConfigData struct {
	RabbitMQConfig RabbitMQConfig `yaml:"RabbitMQConfig"`
	MySQLConfig    MySQLConfig    `yaml:"MySQLConfig"`
	RedisConfig    RedisConfig    `yaml:"RedisConfig"`
}

func InitConfig() error {
	dataBytes, err := os.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println("读取配置失败！" + err.Error())
		return err
	}
	err = yaml.Unmarshal(dataBytes, &configData)
	if err != nil {
		fmt.Println("解析配置失败！" + err.Error())
		return err
	}
	return nil
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
