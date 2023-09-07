package main

import (
	"github.com/awoyai/gin-temp/core"
	"github.com/awoyai/gin-temp/global"
	"github.com/awoyai/gin-temp/initialize"
	"go.uber.org/zap"
)

func main() {
	initialize.OtherInit()
	global.LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.LOG)
	global.DB = initialize.NewGormMysql() // gorm连接数据库
	core.RunWindowsServer()
}
