package middleware

import (
	"fmt"
	"net/http"

	"github.com/awoyai/gin-temp/global"
	"github.com/awoyai/gin-temp/model/common"
	"github.com/awoyai/gin-temp/model/response"
	"github.com/awoyai/gin-temp/repo/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CheckPermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.GetString(common.LoginUserKey)
		path := ctx.Request.URL.Path
		ok, err := system.CasbinRepo.CheckPermission(username, path, "")
		if err != nil {
			global.LOG.Error("CheckPermission fail", zap.Error(err), zap.String("username", username), zap.String("path", path))
			response.FailWithMessage(response.MSG_INTERNAL_ERROR, ctx)
			return
		}
		if !ok {
			global.LOG.Info("CheckPermission not allow", zap.String("username", username), zap.String("path", path))
			response.FailWithUnauthorized(ctx)
			return
		}
	}
}

func PanicHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.LOG.Error("catch server panic", zap.Any("err", err))
				c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("server panic: %v", err))
			}
		}()
		c.Next()
	}
}
