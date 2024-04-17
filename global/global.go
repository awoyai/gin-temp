package global

import (
	"github.com/awoyai/gin-temp/config"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	VP                  *viper.Viper
	DB                  *gorm.DB
	LOG                 *zap.Logger
	CONFIG              *config.Server
	BlackCache          local_cache.Cache
	Concurrency_Control = &singleflight.Group{}
)
