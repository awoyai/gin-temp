package system

import (
	"errors"

	"github.com/awoyai/gin-temp/global"
	"github.com/awoyai/gin-temp/model/common"
	"github.com/awoyai/gin-temp/model/request"
	"github.com/awoyai/gin-temp/model/response"
	"github.com/awoyai/gin-temp/model/system"
	systemRepo "github.com/awoyai/gin-temp/repo/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct{}

func (UserService) Add(ctx *gin.Context) {
	var req request.UserAddReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		global.LOG.Error("UserService#Add#ShouldBind fail", zap.Error(err))
		return
	}
	global.LOG.Info("UserService#Add recv", zap.Any("req", req))
	if req.Username == "" {
		response.FailWithMessage("用户名不能为空", ctx)
		return
	}
	if _, err := systemRepo.UserRepo.Info(map[string]any{"username": req.Username}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user := &system.User{
				NikeName:   req.NikeName,
				Username:   req.Username,
				Department: req.Department,
				Enable:     common.EnableTypeOpen,
				UserType:   req.UserType,
				Roles:      req.Roles,
			}
			if err := systemRepo.UserRepo.Create(user); err != nil {
				response.FailWithMessage(response.MSG_INTERNAL_ERROR, ctx)
				global.LOG.Error("UserService#Add#Create fail", zap.Error(err))
				return
			}
			response.OkWithData(user, ctx)
		} else {
			response.FailWithMessage(response.MSG_INTERNAL_ERROR, ctx)
			global.LOG.Error("UserService#Add#Info fail", zap.Error(err))
		}
		return
	}
	response.FailWithMessage("该用户已存在，请勿重复添加", ctx)
}

func (UserService) Del(ctx *gin.Context) {
	var req request.IDBase
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		global.LOG.Error("UserService#Del#ShouldBind fail", zap.Error(err))
		return
	}
	global.LOG.Info("UserService#Del recv", zap.Any("req", req))
	if err := systemRepo.UserRepo.Delete(req.ID); err != nil {
		response.FailWithMessage(response.MSG_INTERNAL_ERROR, ctx)
		global.LOG.Error("UserService#Del#Delete fail", zap.Error(err))
		return
	}
	response.Ok(ctx)
}

func (UserService) Update(ctx *gin.Context) {
	var req request.UserUpdateReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		global.LOG.Error("UserService#Update#ShouldBind fail", zap.Error(err))
		return
	}
	global.LOG.Info("UserService#Update recv", zap.Any("req", req))
	if err := systemRepo.UserRepo.Update(req.ID, req.ToMap()); err != nil {
		response.FailWithMessage(response.MSG_INTERNAL_ERROR, ctx)
		global.LOG.Error("UserService#Update#Update fail", zap.Error(err))
		return
	}
	response.Ok(ctx)
}

func (UserService) List(ctx *gin.Context) {
	var req request.UserListReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		global.LOG.Error("UserService#List#ShouldBind fail", zap.Error(err))
		return
	}
	global.LOG.Info("UserService#List recv", zap.Any("req", req))
	filter := &system.UserFilter{
		Username: req.Username,
		PageInfo: req.PageInfo,
	}
	list, err := systemRepo.UserRepo.List(filter)
	if err != nil {
		response.FailWithMessage(response.MSG_INTERNAL_ERROR, ctx)
		global.LOG.Error("UserService#List#List fail", zap.Error(err))
		return
	}
	response.OkWithData(common.PageResult{
		List:     list,
		PageInfo: filter.PageInfo,
	}, ctx)
}
