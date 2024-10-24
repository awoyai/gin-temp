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

type MenuService struct{}

func (MenuService) GetAllPaths(ctx *gin.Context) {
	var AllPaths []system.PathInfo
	for _, router := range global.G.Routes() {
		AllPaths = append(AllPaths, system.PathInfo{Path: router.Path, Method: router.Method})
	}
	response.OkWithData(AllPaths, ctx)
}

func (MenuService) Add(ctx *gin.Context) {
	var req system.Menu
	if err := ctx.ShouldBind(&req); err != nil {
		global.LOG.Error("AddMenu#ShouldBind", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if req.Name == "" || req.Key == "" {
		response.FailWithMessage(response.MSG_BAD_PARAM, ctx)
		return
	}
	req.ID = 0
	global.LOG.Info("AddMenu#recv", zap.Any("req", req))
	_, err := systemRepo.MenuRepo.InfoByWr(map[string]any{"key": req.Key})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.LOG.Error("Add#InfoByWr", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	} else if err == nil {
		response.FailWithMessage("菜单Key不可重复", ctx)
		return
	}
	if err := systemRepo.MenuRepo.Add(&req); err != nil {
		global.LOG.Error("AddMenu#Add", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}

func (MenuService) Update(ctx *gin.Context) {
	var req system.Menu
	if err := ctx.ShouldBind(&req); err != nil {
		global.LOG.Error("Update#ShouldBind", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	global.LOG.Info("Update#recv", zap.Any("req", req))
	if err := systemRepo.MenuRepo.Update(req.ID, map[string]any{
		"name":      req.Name,
		"pid":       req.Pid,
		"path":      req.Path,
		"redirect":  req.Redirect,
		"component": req.Component,
		"icon":      req.Icon,
		"sort":      req.Sort,
		"api_list":  req.APIList,
		"menu_type": req.MenuType,
	}); err != nil {
		global.LOG.Error("Update#Update", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}

func (MenuService) Delete(ctx *gin.Context) {
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
	if err := systemRepo.MenuRepo.Delete(req.ID); err != nil {
		global.LOG.Error("Delete#Delete", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}

func (MenuService) ALLMenuList(ctx *gin.Context) {
	menuList, err := systemRepo.MenuRepo.AllList()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(menuList, ctx)
}

func (MenuService) AuthMenuList(ctx *gin.Context) {
	username := ctx.GetString(common.LoginUserKey)
	menuList, err := systemRepo.RoleRepo.AuthMenuList(username)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(systemRepo.MenuRepo.ConstructTree(menuList), ctx)
}
