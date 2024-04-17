package initialize

import (
	"github.com/awoyai/gin-temp/middleware"
	"github.com/awoyai/gin-temp/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	g := gin.Default()
	publicGroup := g.Group("/api/v1")
	router.InitBaseRouter(publicGroup)
	privateGroup := g.Group("/api/v1", middleware.JWTAuth())
	router.InitGreeterRouter(privateGroup)
	return g
}
