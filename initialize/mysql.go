package initialize

import (
	"log"
	"os"
	"time"

	"github.com/awoyai/gin-temp/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewGormMysql() *gorm.DB {
	m := global.CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	l := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	l.LogMode(logger.Error)
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   global.CONFIG.Mysql.Prefix,
			SingularTable: global.CONFIG.Mysql.Singular,
		},
		Logger:                                   l,
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
