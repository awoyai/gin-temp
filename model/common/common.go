package common

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

var (
	LoginUserKey = "username"
)

type EnableType int

const (
	EnableTypeOpen = iota + 1
	EnableTypeClose
)

type PageResult struct {
	List any `json:"list"`
	PageInfo
}

type PageInfo struct {
	PageNum    int64 `json:"page_num"`
	PageSize   int64 `json:"page_size"` // 每页大小
	TotalCount int64 `json:"total_count"`
	TotalPage  int64 `json:"total_page"`
}

type OrderBy struct {
	OrderField     string `json:"order_field"`
	OrderDirection string `json:"order_direction"`
}

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UintSlice []uint

func (m *UintSlice) Scan(val interface{}) error {
	s, _ := val.([]byte)
	return json.Unmarshal(s, m)
}

func (m UintSlice) Value() (driver.Value, error) {
	return json.Marshal(m)
}

type StringSlice []string

func (m *StringSlice) Scan(val interface{}) error {
	s, _ := val.([]byte)
	return json.Unmarshal(s, m)
}

func (m StringSlice) Value() (driver.Value, error) {
	return json.Marshal(m)
}
