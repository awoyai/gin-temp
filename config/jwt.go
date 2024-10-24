package config

type JWT struct {
	Secret  string `mapstructure:"secret" json:"secret" yaml:"secret"`       // jwt签名
	TimeOut int    `mapstructure:"time_out" json:"time_out" yaml:"time_out"` // 过期时间
}
