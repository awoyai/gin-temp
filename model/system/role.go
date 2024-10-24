package system

import (
	"github.com/awoyai/gin-temp/model/common"
	"github.com/awoyai/gin-temp/model/response"
)

const (
	DOMAIN_DEFAULT = "main"
)

type ListRole []*Role

type Role struct {
	common.MODEL
	Name        string           `gorm:"cloumn:name" json:"name"`
	Code        string           `gorm:"cloumn:code" json:"code"`
	Auths       common.UintSlice `gorm:"cloumn:auths;type:varchar(10000)" json:"auths"`
	Description string           `gorm:"cloumn:description" json:"description"`
}

func (r Role) GetCode() string {
	return "Role::" + r.Code
}

type RoleFilter struct {
	Name string
	common.PageInfo
}

func (Role) TableName() string {
	return "tb_roles"
}

func (l ListRole) ToRsp() response.ListRoleRsp {
	res := make(response.ListRoleRsp, len(l))
	for i, v := range l {
		res[i] = v.ToRsp()
	}
	return res
}

func (r Role) ToRsp() *response.RoleRsp {
	return &response.RoleRsp{
		ID:   r.ID,
		Name: r.Name,
	}
}

type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:100;uniqueIndex:unique_index"`
	V0    string `gorm:"size:100;uniqueIndex:unique_index"`
	V1    string `gorm:"size:100;uniqueIndex:unique_index"`
	V2    string `gorm:"size:100;uniqueIndex:unique_index"`
	V3    string `gorm:"size:100;uniqueIndex:unique_index"`
	V4    string `gorm:"size:100;uniqueIndex:unique_index"`
	V5    string `gorm:"size:100;uniqueIndex:unique_index"`
}

func (CasbinRule) TableName() string {
	return "tb_casbin"
}
