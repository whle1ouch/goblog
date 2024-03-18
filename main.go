package main

import (
	"goblog/core"
	"goblog/global"
)

func main() {
	// 配置初始化
	core.InitConfig()

	// 初始化日志
	global.Log = core.InitLogger()
	global.Log.Info("日志初始化成功")
	global.Log.Error("sss")
	// 连接数据库
	// global.DB = core.InitGorm()
}
