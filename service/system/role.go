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

type RoleService struct{}

func (RoleService) Add(ctx *gin.Context) {
	var req system.Role
	if err := ctx.ShouldBind(&req); err != nil {
		global.LOG.Error("Add#ShouldBind", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if req.Name == "" {
		response.FailWithMessage(response.MSG_BAD_PARAM, ctx)
		return
	}
	req.ID = 0
	global.LOG.Info("Add#recv", zap.Any("req", req))
	_, err := systemRepo.RoleRepo.InfoByWr(map[string]any{"code": req.Code})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.LOG.Error("Add#InfoByWr", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	} else if err == nil {
		response.FailWithMessage("角色Code不可重复", ctx)
		return
	}
	if err := systemRepo.RoleRepo.Add(&req); err != nil {
		global.LOG.Error("Add#Add", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}

func (RoleService) Update(ctx *gin.Context) {
	var req system.Role
	if err := ctx.ShouldBind(&req); err != nil {
		global.LOG.Error("Update#ShouldBind", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	global.LOG.Info("Update#recv", zap.Any("req", req))
	if err := systemRepo.RoleRepo.Update(req.ID, map[string]any{
		"name":        req.Name,
		"auths":       req.Auths,
		"description": req.Description,
	}); err != nil {
		global.LOG.Error("Update#Update", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}

func (RoleService) Delete(ctx *gin.Context) {
	var req request.IDBase
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if req.ID == 0 {
		response.FailWithMessage(response.MSG_BAD_PARAM, ctx)
		return
	}
	global.LOG.Info("Delete#recv", zap.Any("req", req))
	if err := systemRepo.RoleRepo.Delete(req.ID); err != nil {
		global.LOG.Error("Delete#Delete", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}

func (RoleService) List(ctx *gin.Context) {
	var req request.BaseListReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	global.LOG.Info("List#recv", zap.Any("req", req))
	filter := system.RoleFilter{Name: req.Name, PageInfo: req.PageInfo}
	list, err := systemRepo.RoleRepo.List(&filter)
	if err != nil {
		global.LOG.Error("Delete#Delete", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(common.PageResult{
		List:     list,
		PageInfo: filter.PageInfo,
	}, ctx)
}
