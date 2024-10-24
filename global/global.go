package global

import (
	"github.com/awoyai/gin-temp/config"
	"github.com/awoyai/gin-temp/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	G      *gin.Engine = gin.Default()
	VP     *viper.Viper
	DB     *gorm.DB
	DBMap  map[string]*gorm.DB
	LOG    *zap.Logger
	CONFIG *config.Server
	JWT    *utils.JWTMiddleware
)

func GetDB() *gorm.DB {
	return DB
}

// func GetSpiderDB() *gorm.DB {
// 	return DBMap["spider"]
// }
