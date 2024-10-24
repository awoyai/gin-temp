package router

import (
	"github.com/awoyai/gin-temp/middleware"
	"github.com/awoyai/gin-temp/service"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(router *gin.RouterGroup) {
	baseRouter := router.Group("")
	baseSrv := service.ServiceGroupApp.SystemServiceGroup.BaseSrv
	{
		baseRouter.POST("login", baseSrv.Login)
		baseRouter.POST("unsafe_login", baseSrv.UnsafeLogin)
	}
}

func InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")
	userSrv := service.ServiceGroupApp.SystemServiceGroup.UserSrv
	{
		userRouter.POST("add", userSrv.Add)
		userRouter.POST("del", userSrv.Del)
		userRouter.POST("update", userSrv.Update)
		userRouter.POST("list", userSrv.List)
	}
}

func InitRoleRouter(router *gin.RouterGroup) {
	roleRouter := router.Group("role")
	roleSrv := service.ServiceGroupApp.SystemServiceGroup.RoleSrv
	{
		roleRouter.POST("add", roleSrv.Add)
		roleRouter.POST("del", roleSrv.Delete)
		roleRouter.POST("update", roleSrv.Update)
		roleRouter.POST("list", roleSrv.List)
	}
}

func InitMenuRouter(router *gin.RouterGroup) {
	menuRouter := router.Group("menu")
	menuSrv := service.ServiceGroupApp.SystemServiceGroup.MenuSrv
	{
		menuRouter.POST("add", menuSrv.Add)
		menuRouter.POST("del", menuSrv.Delete)
		menuRouter.POST("update", menuSrv.Update)
		menuRouter.GET("list", middleware.CheckPermission(), menuSrv.ALLMenuList)
		menuRouter.GET("api/list", menuSrv.GetAllPaths)
		menuRouter.GET("list/auth", menuSrv.AuthMenuList)
	}
}
