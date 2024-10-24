package initialize

import (
	"github.com/awoyai/gin-temp/global"
	"github.com/awoyai/gin-temp/middleware"
	"github.com/awoyai/gin-temp/router"
	"github.com/awoyai/gin-temp/service"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	global.G.Use(gin.Recovery(), middleware.PanicHandle())
	global.G.GET("/health", service.ServiceGroupApp.SystemServiceGroup.BaseSrv.HealthyCheck)
	publicGroup := global.G.Group(global.CONFIG.System.RouterPrefix)
	router.InitBaseRouter(publicGroup)
	privateGroup := global.G.Group(global.CONFIG.System.RouterPrefix, global.JWT.MiddlewareFunc())
	{
		// system
		router.InitUserRouter(privateGroup)
		router.InitMenuRouter(privateGroup)
		router.InitRoleRouter(privateGroup)
	}

	{
		// greeter
		router.InitGreeterRouter(privateGroup)
	}
	return global.G
}
