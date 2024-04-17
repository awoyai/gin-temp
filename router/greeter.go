package router

import (
	"github.com/awoyai/gin-temp/service"
	"github.com/gin-gonic/gin"
)

func InitGreeterRouter(router *gin.RouterGroup) {
	greeterRouter := router.Group("greeter")
	greeterSrv := service.ServiceGroupApp.GreeterSrv
	{
		greeterRouter.POST("greeter", greeterSrv.Greeter)
	}
}
