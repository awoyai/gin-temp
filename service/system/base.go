package system

import (
	"errors"

	"github.com/awoyai/gin-temp/global"
	"github.com/awoyai/gin-temp/model/common"
	"github.com/awoyai/gin-temp/model/request"
	"github.com/awoyai/gin-temp/model/response"
	systemRepo "github.com/awoyai/gin-temp/repo/system"
	"github.com/awoyai/gin-temp/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BaseService struct{}

func (BaseService) HealthyCheck(ctx *gin.Context) {
	response.OkWithMessage("success", ctx)
}

func (BaseService) Login(ctx *gin.Context) {
	var req request.LoginReq
	info, err := systemRepo.UserRepo.Info(map[string]any{"username": req.Username})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.FailWithMessage("未有登录权限，请联系管理员添加", ctx)
			return
		}
		response.FailWithMessage(response.MSG_INTERNAL_ERROR, ctx)
		global.LOG.Error("Login#Login", zap.Error(err))
		return
	}
	if info.Enable != common.EnableTypeOpen {
		response.FailForbidden(ctx)
		return
	}
	token, expired, err := global.JWT.TokenGenerator(map[string]interface{}{utils.TOKEN_KEY_NAME: req.Username})
	if err != nil {
		response.FailWithMessage(response.MSG_INTERNAL_ERROR, ctx)
		global.LOG.Error("Login#TokenGenerator", zap.Error(err))
		return
	}
	menuList, err := systemRepo.RoleRepo.AuthMenuList(info.Username)
	if err != nil {
		response.FailWithMessage(response.MSG_INTERNAL_ERROR, ctx)
		global.LOG.Error("Login#AuthMenuList", zap.Error(err))
		return
	}
	permissions := make([]string, 0)
	for _, menu := range menuList {
		permissions = append(permissions, menu.Key)
	}
	global.JWT.SendJWTCookie(ctx, token, expired)
	response.OkWithData(map[string]any{"username": info.Username, "permission": permissions}, ctx)
}

func (BaseService) UnsafeLogin(ctx *gin.Context) {
	username := ctx.Request.Header.Get("username")
	info, err := systemRepo.UserRepo.Info(map[string]any{"username": username})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.FailWithUnauthorized(ctx)
			return
		}
		response.FailWithMessage(response.MSG_INTERNAL_ERROR, ctx)
		global.LOG.Error("UnsafeLogin#UserRepo#Info", zap.Error(err))
		return
	}
	token, expired, err := global.JWT.TokenGenerator(map[string]interface{}{utils.TOKEN_KEY_NAME: username})
	if err != nil {
		response.FailWithMessage(response.MSG_INTERNAL_ERROR, ctx)
		global.LOG.Error("Login#JWT#TokenGenerator", zap.Error(err))
		return
	}
	menuList, err := systemRepo.RoleRepo.AuthMenuList(info.Username)
	if err != nil {
		response.FailWithMessage(response.MSG_INTERNAL_ERROR, ctx)
		global.LOG.Error("Login#AuthMenuList", zap.Error(err))
		return
	}
	permissions := make([]string, 0)
	for _, menu := range menuList {
		permissions = append(permissions, menu.Key)
	}
	global.JWT.SendJWTCookie(ctx, token, expired)
	response.OkWithData(map[string]any{"username": info.Username, "permission": permissions}, ctx)
}
