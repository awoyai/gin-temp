package request

import (
	"github.com/awoyai/gin-temp/model/common"
)

type IDBase struct {
	ID uint `json:"id"`
}

type BaseListReq struct {
	Name string `json:"name"`
	common.PageInfo
}
