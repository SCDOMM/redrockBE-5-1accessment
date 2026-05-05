package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var configData ConfigData

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
	HertzConfig HertzConfig `yaml:"HertzConfig"`
	KitexConfig KitexConfig `yaml:"KitexConfig"`
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
func GetKitexConfig() KitexConfig {
	return configData.KitexConfig
}
func GetHertzConfig() HertzConfig {
	return configData.HertzConfig
}
