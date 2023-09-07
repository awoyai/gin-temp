package router

import (
	"github.com/awoyai/gin-temp/service"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(router *gin.RouterGroup) {
	baseRouter := router.Group("base")
	baseSrv := service.ServiceGroupApp.BaseSrv
	{
		baseRouter.POST("login", baseSrv.Login)
		baseRouter.POST("captcha", baseSrv.Captcha)
	}
}
