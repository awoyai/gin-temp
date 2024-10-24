package config

type Server struct {
	Zap         Zap         `mapstructure:"zap" json:"zap" yaml:"zap"`
	JWT         JWT         `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Mysql       Mysql       `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	System      System      `mapstructure:"system" json:"system" yaml:"system"`
}