package system

import (
	"github.com/awoyai/gin-temp/model/common"
)

type MenuType int

const (
	MenuTypeMenu = iota
	MenuTypeButton
)

const (
	MenuTypeMenuStr   = "menu"
	MenuTypeButtonStr = "button"
)

func NewMenuType(typ string) MenuType {
	switch typ {
	case MenuTypeButtonStr:
		return MenuTypeButton
	default:
		return MenuTypeMenu
	}
}

type Menu struct {
	common.MODEL
	Name      string             `json:"name"`
	Pid       uint               `json:"pid"`
	Path      string             `json:"path"`
	Component string             `json:"component"`
	Redirect  string             `json:"redirect"`
	Icon      string             `json:"icon"`
	Sort      int                `json:"sort"`
	MenuType  MenuType           `json:"menu_type"`
	Key       string             `json:"key"`
	APIList   common.StringSlice `json:"api_list" gorm:"column:api_list"`
	Children  []*Menu            `json:"children" gorm:"-"`
	Auth
}

func (m Menu) GetKey() string {
	return "Menu::" + m.Key
}

func (Menu) TableName() string {
	return "tb_menu"
}

type MenuFilter struct {
	IDList []uint
	common.PageInfo
}

type PathInfo struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}
