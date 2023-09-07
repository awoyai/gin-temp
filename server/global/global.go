package global

import (
	"github.com/awoyai/gin-temp/config"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	DB                  *gorm.DB
	LOG                 *zap.Logger
	CONFIG              *config.Server
	BlackCache          local_cache.Cache
	Concurrency_Control = &singleflight.Group{}
)
