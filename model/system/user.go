package system

import (
	"github.com/awoyai/gin-temp/model/common"
)

type UserType int

const (
	UserTypeAdmin = iota + 1
	UserTypeNormal
)

type User struct {
	common.MODEL
	NikeName   string            `gorm:"column:nike_name;type:varchar(255);comment:名字" json:"nike_name"`
	Username   string            `gorm:"column:username;type:varchar(255);comment:用户名" json:"username"`
	Department string            `gorm:"column:department;type:varchar(255);comment:部门" json:"department"`
	Roles      common.UintSlice  `gorm:"column:roles;type:varchar(10000);comment:权限" json:"roles"`
	Enable     common.EnableType `gorm:"column:enable;type:tinyint(1);comment:是否启用" json:"enable"`
	UserType   UserType          `gorm:"column:user_type;type:tinyint(1);comment:用户类型" json:"user_type"`
}

func (User) TableName() string {
	return "tb_user"
}

type UserFilter struct {
	ID       uint
	Username string
	common.PageInfo
}
