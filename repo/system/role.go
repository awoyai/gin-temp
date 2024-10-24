package system

import (
	"github.com/awoyai/gin-temp/global"
	"github.com/awoyai/gin-temp/model/common"
	"github.com/awoyai/gin-temp/model/system"
	"github.com/awoyai/gin-temp/utils"
	"gorm.io/gorm"
)

var RoleRepo = new(roleRepo)

type roleRepo struct{}

func (r roleRepo) Add(info *system.Role) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(info).Error; err != nil {
			return err
		}
		if err := r.addPolice(tx, info.GetCode(), info.Auths); err != nil {
			return err
		}
		return nil
	})
}

func (roleRepo) addPolice(tx *gorm.DB, roleCode string, menuIDList []uint) error {
	var menuList []*system.Menu
	if len(menuIDList) != 0 {
		if err := tx.Where("id in ?", menuIDList).Find(&menuList).Error; err != nil {
			return err
		}
	}
	menus := []string{}
	if len(menuList) != 0 {
		for _, menu := range menuList {
			menus = append(menus, menu.GetKey())
		}
		if _, err := CasbinRepo.Casbin().AddRolesForUser(roleCode, menus); err != nil {
			return err
		}
	}
	return nil
}

func (r roleRepo) Update(id uint, data map[string]any) error {
	role, err := r.Info(id)
	if err != nil {
		return err
	}
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&system.Role{}).Where("id = ?", id).Updates(data).Error; err != nil {
			return err
		}
		if auths, ok := data["auths"]; ok {
			if eq := utils.UintSliceEq(auths.(common.UintSlice), role.Auths); !eq {
				if _, err := CasbinRepo.Casbin().RemoveFilteredPolicy(0, role.GetCode()); err != nil {
					return err
				}
				if err := r.addPolice(tx, role.GetCode(), auths.(common.UintSlice)); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (roleRepo) Delete(id uint) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		var role system.Role
		if err := tx.First(&role, id).Error; err != nil {
			return err
		}
		if err := tx.Delete(&system.Role{}, id).Error; err != nil {
			return err
		}
		if _, err := CasbinRepo.Casbin().DeleteRole(role.Code); err != nil {
			return err
		}
		return nil
	})
}

func (roleRepo) List(filter *system.RoleFilter) ([]*system.Role, error) {
	var res []*system.Role
	db := global.DB.Model(&system.Role{})
	if filter.Name != "" {
		db = db.Where("name = ?", filter.Name)
	}
	if filter.PageNum != 0 && filter.PageSize != 0 {
		db = db.Count(&filter.TotalCount).Offset(int((filter.PageNum - 1) * filter.PageSize)).Limit(int(filter.PageSize))
	}
	if err := db.Order("name asc").Find(&res).Error; err != nil {
		return nil, err
	}
	if filter.PageNum != 0 && filter.PageSize != 0 {
		filter.TotalPage = (filter.TotalCount + filter.PageSize - 1) / filter.PageSize
	}
	return res, nil
}

func (roleRepo) Info(id uint) (*system.Role, error) {
	var role *system.Role
	return role, global.DB.Last(&role, id).Error
}

func (roleRepo) InfoByWr(wr map[string]any) (*system.Role, error) {
	var info system.Role
	return &info, global.GetDB().Model(&system.Role{}).Where(wr).First(&info).Error
}

func (roleRepo) AuthMenuList(username string) ([]*system.Menu, error) {
	var user system.User
	if err := global.DB.Where("username = ? and enable = ?", username, common.EnableTypeOpen).First(&user).Error; err != nil {
		return nil, err
	}
	var roleList []*system.Role
	if len(user.Roles) != 0 {
		if err := global.DB.Debug().Where("id in ?", []uint(user.Roles)).Find(&roleList).Error; err != nil {
			return nil, err
		}
	}
	menuIDList := make([]uint, 0)
	for _, role := range roleList {
		menuIDList = append(menuIDList, role.Auths...)
	}
	var menuList []*system.Menu
	if len(menuIDList) != 0 {
		if err := global.DB.Where("id in ?", []uint(menuIDList)).Find(&menuList).Error; err != nil {
			return nil, err
		}
	}
	return menuList, nil
}
