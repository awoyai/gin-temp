package global

import (
	"time"
)

type MODEL struct {
	ID        uint      `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}
