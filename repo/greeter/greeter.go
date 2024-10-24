package spider

import (
	"github.com/awoyai/gin-temp/global"
	"github.com/awoyai/gin-temp/model/greeter"
)

var GreeterRepo = new(greeterRepo)

type greeterRepo struct{}

func (greeterRepo) Create(info *greeter.Greeter) error {
	return global.GetDB().Where("id = ?", info.ID).FirstOrCreate(&info).Error
}

func (greeterRepo) Update(id uint, data map[string]any) error {
	return global.GetDB().Model(&greeter.Greeter{}).Where("id = ?", id).Updates(data).Error
}

func (greeterRepo) Delete(id uint) error {
	return global.GetDB().Delete(&greeter.Greeter{}, id).Error
}

func (greeterRepo) Info(wr map[string]any) (*greeter.Greeter, error) {
	var res greeter.Greeter
	return &res, global.GetDB().Model(&greeter.Greeter{}).Where(wr).First(&res).Error
}

func (greeterRepo) List(filter *greeter.GreeterFilter) ([]*greeter.Greeter, error) {
	var res []*greeter.Greeter
	db := global.GetDB().Model(&greeter.Greeter{})
	if filter.Name != "" {
		db = db.Where("name like ?", "%"+filter.Name+"%")
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
