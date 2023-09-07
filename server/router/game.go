package router

import (
	"github.com/awoyai/gin-temp/service"
	"github.com/gin-gonic/gin"
)

func InitGameRouter(router *gin.RouterGroup) {
	gameRouter := router.Group("game")
	gameSrv := service.ServiceGroupApp.GameSrv
	{
		gameRouter.POST("add", gameSrv.CreateGame)
	}
}
