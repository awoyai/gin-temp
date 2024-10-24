package greeter

import (
	"github.com/awoyai/gin-temp/global"
	"github.com/awoyai/gin-temp/model/common"
	"github.com/awoyai/gin-temp/model/greeter"
	"github.com/awoyai/gin-temp/model/request"
	"github.com/awoyai/gin-temp/model/response"
	geeterRepo "github.com/awoyai/gin-temp/repo/greeter"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GreeterService struct{}

func (GreeterService) List(ctx *gin.Context) {
	var req request.GreeterReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		global.LOG.Error("UserService#List#ShouldBind fail", zap.Error(err))
		return
	}
	global.LOG.Info("UserService#List recv", zap.Any("req", req))
	filter := greeter.GreeterFilter{
		Name:     req.Name,
		PageInfo: req.PageInfo,
	}
	list, err := geeterRepo.GreeterRepo.List(&filter)
	if err != nil {
		response.FailWithMessage(response.MSG_SCHEME_SEARCH_ERROR, ctx)
		global.LOG.Error("GreeterService#List#List fail", zap.Error(err))
		return
	}
	response.OkWithData(common.PageResult{List: list, PageInfo: filter.PageInfo}, ctx)
}
