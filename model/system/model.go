package system

type Auth struct {
	Permission bool `gorm:"-" json:"permission"`
}
