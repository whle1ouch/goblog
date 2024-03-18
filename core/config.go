package core

import (
	"fmt"
	"goblog/config"
	"goblog/global"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// 读取配置文件
func InitConfig() {
	const ConfigFile = "setting.yaml"
	c := &config.Config{}
	yamlConfif, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败：%v", err))
	}
	err = yaml.Unmarshal(yamlConfif, c)
	if err != nil {
		log.Fatalf("解析配置文件失败：%v", err)
	}
	log.Println("配置文件读取成功")
	global.Config = c
}
