package core

import (
	"fmt"
	"goblog/config"
	"goblog/global"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const ConfigFile = "setting.yaml"

// 读取配置文件
func InitConfig() {

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

// 把配置写入yaml文件
func WriteConfigToYaml() (err error) {
	yamlConfig, err := yaml.Marshal(global.Config)
	if err != nil {
		log.Fatalf("写入配置文件失败：%v", err)
		return
	}
	err = os.WriteFile(ConfigFile, yamlConfig, 0666)
	if err != nil {
		log.Fatalf("写入配置文件失败：%v", err)
		return
	}
	log.Println("写入配置文件成功")
	return nil
}
