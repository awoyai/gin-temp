package request

import "github.com/awoyai/gin-temp/model/common"

type GreeterReq struct {
	Name string `json:"name"`
	common.PageInfo
}
