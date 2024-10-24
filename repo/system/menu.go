package system

import (
	"github.com/awoyai/gin-temp/global"
	"github.com/awoyai/gin-temp/model/common"
	"github.com/awoyai/gin-temp/model/system"
	"github.com/awoyai/gin-temp/utils"
	"gorm.io/gorm"
)

var MenuRepo = new(menuRepo)

type menuRepo struct{}

func (r menuRepo) Add(info *system.Menu) error {
	return global.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := r.addPolice(info.GetKey(), info.APIList); err != nil {
			return err
		}
		if err := tx.Create(info).Error; err != nil {
			return err
		}
		return nil
	})
}
func (menuRepo) addPolice(menuKey string, apiList []string) error {
	rules := [][]string{}
	if len(apiList) != 0 {
		for _, api := range apiList {
			rules = append(rules, []string{menuKey, api})
		}
		if _, err := CasbinRepo.Casbin().AddPolicies(rules); err != nil {
			return err
		}
	}
	return nil
}

func (r menuRepo) Update(id uint, data map[string]any) error {
	menu, err := r.Info(id)
	if err != nil {
		return err
	}
	return global.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&system.Menu{}).Where("id = ?", id).Updates(data).Error; err != nil {
			return err
		}
		if apiList, ok := data["api_list"]; ok {
			if eq := utils.StrignSliceEq(apiList.(common.StringSlice), menu.APIList); !eq {
				if _, err := CasbinRepo.Casbin().RemoveFilteredPolicy(0, menu.GetKey()); err != nil {
					return err
				}
				if err := r.addPolice(menu.GetKey(), apiList.(common.StringSlice)); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (menuRepo) Delete(id uint) error {
	return global.GetDB().Delete(&system.Menu{}, id).Error
}

func (menuRepo) ListByFilter(filter *system.MenuFilter) ([]*system.Menu, error) {
	var list []*system.Menu
	db := global.DB.Model(&system.Menu{})
	if len(filter.IDList) != 0 {
		db = db.Where("id in ?", filter.IDList)
	}
	return list, db.Find(&list).Error
}

func (r menuRepo) AllList() ([]*system.Menu, error) {
	var list []*system.Menu
	if err := global.DB.Order("sort ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	return r.ConstructTree(list), nil
}

func (r menuRepo) TreeList(pid []uint) (map[uint][]*system.Menu, error) {
	var menuList []*system.Menu
	if err := global.GetDB().Model(&system.Menu{}).Where("pid in ?", pid).Order("sort ASC").Find(&menuList).Error; err != nil {
		return nil, err
	}
	if len(menuList) == 0 {
		return nil, nil
	}
	m := make(map[uint][]*system.Menu)
	pidList := make([]uint, len(menuList))
	for i, v := range menuList {
		pidList[i] = v.ID
		m[v.Pid] = append(m[v.Pid], v)
	}
	childrenMap, err := r.TreeList(pidList)
	if err != nil {
		return nil, err
	}
	if childrenMap != nil {
		for _, v := range menuList {
			v.Children = childrenMap[v.ID]
		}
	}
	return m, nil
}

func (menuRepo) ConstructTree(menuList []*system.Menu) []*system.Menu {
	m := make(map[uint]*system.Menu)
	treeList := make([]*system.Menu, 0)
	for _, v := range menuList {
		m[v.ID] = v
	}
	for _, v := range menuList {
		if v.Pid == 0 {
			treeList = append(treeList, v)
		} else {
			p, ok := m[v.Pid]
			if ok {
				p.Children = append(p.Children, v)
			}
		}
	}
	return treeList
}

func (menuRepo) Info(id uint) (*system.Menu, error) {
	var menu *system.Menu
	return menu, global.GetDB().Last(&menu, id).Error
}

func (menuRepo) InfoByWr(wr map[string]any) (*system.Menu, error) {
	var info system.Menu
	return &info, global.GetDB().Model(&system.Menu{}).Where(wr).First(&info).Error
}
