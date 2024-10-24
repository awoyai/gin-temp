package system

import (
	"github.com/awoyai/gin-temp/global"
	"github.com/awoyai/gin-temp/model/common"
	"github.com/awoyai/gin-temp/model/system"
	"github.com/awoyai/gin-temp/utils"
	"gorm.io/gorm"
)

var UserRepo = new(userRepo)

type userRepo struct{}

func (r userRepo) Create(info *system.User) error {
	return global.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := r.addPolice(tx, info.Username, info.Roles); err != nil {
			return err
		}
		return tx.Where("id = ?", info.ID).FirstOrCreate(&info).Error
	})
}

func (userRepo) addPolice(tx *gorm.DB, username string, roleIDList []uint) error {
	var roleList []*system.Role
	if len(roleIDList) != 0 {
		if err := tx.Where("id in ?", roleIDList).Find(&roleList).Error; err != nil {
			return err
		}
	}
	roles := []string{}
	if len(roleList) != 0 {
		for _, role := range roleList {
			roles = append(roles, role.GetCode())
		}
		if _, err := CasbinRepo.Casbin().AddRolesForUser(username, roles); err != nil {
			return err
		}
	}
	return nil
}

func (r userRepo) Update(id uint, data map[string]any) error {
	user, err := r.Info(map[string]any{"id": id})
	if err != nil {
		return err
	}
	return global.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&system.User{}).Where("id = ?", id).Updates(data).Error; err != nil {
			return err
		}
		if roles, ok := data["roles"]; ok {
			if eq := utils.UintSliceEq(roles.(common.UintSlice), user.Roles); !eq {
				if _, err := CasbinRepo.Casbin().DeleteRolesForUser(user.Username); err != nil {
					return err
				}
				if err := r.addPolice(tx, user.Username, roles.(common.UintSlice)); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (userRepo) Delete(id uint) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		var user system.User
		if err := tx.First(&user, id).Error; err != nil {
			return err
		}
		if err := tx.Delete(&system.User{}, id).Error; err != nil {
			return err
		}
		if _, err := CasbinRepo.Casbin().DeleteUser(user.Username); err != nil {
			return err
		}
		return nil
	})
}

func (userRepo) Info(wr map[string]any) (*system.User, error) {
	var res system.User
	return &res, global.GetDB().Model(&system.User{}).Where(wr).First(&res).Error
}

func (userRepo) List(filter *system.UserFilter) ([]*system.User, error) {
	var res []*system.User
	db := global.GetDB().Model(&system.User{})
	if filter.ID != 0 {
		db = db.Where("id = ?", filter.ID)
	}
	if filter.Username != "" {
		db = db.Where("username like ?", "%"+filter.Username+"%")
	}
	if filter.PageNum != 0 && filter.PageSize != 0 {
		db = db.Count(&filter.TotalCount).Offset(int((filter.PageNum - 1) * filter.PageSize)).Limit(int(filter.PageSize))
	}
	if err := db.Find(&res).Error; err != nil {
		return nil, err
	}
	if filter.PageNum != 0 && filter.PageSize != 0 {
		filter.TotalPage = (filter.TotalCount + filter.PageSize - 1) / filter.PageSize
	}
	return res, nil
}
