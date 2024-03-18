package main

import (
	"goblog/core"
	"goblog/global"
	"goblog/router"
)

func main() {
	// 配置初始化
	core.InitConfig()

	// 初始化日志
	global.Log = core.InitLogger()
	global.Log.Infof("日志初始化成功")
	global.Log.Errorf("sss")
	// 连接数据库
	global.DB = core.InitGorm()
	if global.DB == nil {
		return
	}

	router := router.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("go blog项目启动服务: %s", addr)
	router.Run()
}
