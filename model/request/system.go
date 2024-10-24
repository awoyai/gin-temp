package request

import (
	"github.com/awoyai/gin-temp/model/common"
	"github.com/awoyai/gin-temp/model/system"
)

type (
	LoginReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	UserAddReq struct {
		NikeName   string           `json:"nike_name"`
		Username   string           `json:"username"`
		Department string           `json:"department"`
		UserType   system.UserType  `json:"user_type"`
		Roles      common.UintSlice `json:"roles"`
	}

	UserUpdateReq struct {
		IDBase
		NikeName   string            `json:"nike_name"`
		Department string            `json:"department"`
		UserType   system.UserType   `json:"user_type"`
		Enable     common.EnableType `json:"enable"`
		Roles      common.UintSlice  `json:"roles"`
	}

	UserListReq struct {
		Username string `json:"username"`
		common.PageInfo
	}
)

func (req UserUpdateReq) ToMap() map[string]any {
	return map[string]any{
		"nike_name":  req.NikeName,
		"department": req.Department,
		"user_type":  req.UserType,
		"enable":     req.Enable,
		"roles":      req.Roles,
	}
}
