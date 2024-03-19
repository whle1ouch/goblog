package main

import (
	"goblog/core"
	"goblog/flags"
	"goblog/global"
	"goblog/router"
)

func main() {
	// 配置初始化
	core.InitConfig()
	option := flags.Parse()
	option.ShowVersion()

	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
	if global.DB == nil {
		return
	}

	// 数据库迁移

	if option.IsServerStop() {
		option.Migration()
		return
	}

	// 初始化路由
	router := router.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("go blog项目启动服务: %s", addr)
	router.Run()
}
