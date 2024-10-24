package greeter

import "github.com/awoyai/gin-temp/model/common"

type Greeter struct {
	common.MODEL
}

type GreeterFilter struct {
	Name string
    common.PageInfo
}
