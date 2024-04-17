package core

import (
	"fmt"
	"time"

	"github.com/awoyai/gin-temp/global"
	"github.com/awoyai/gin-temp/initialize"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.CONFIG.System.Port)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	网站已启动: 127.0.0.1:%s
	`, address)
	global.LOG.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
