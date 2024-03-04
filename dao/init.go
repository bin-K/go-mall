package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

var _db *gorm.DB

func DataBase(connRead string, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		// 打印日志信息
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize:         256,  // string类型默认长度
		DisableDatetimePrecision:  true, // 禁止时间精度， 5.6前不支持
		DontSupportRenameColumn:   true, // 用change重命名列， 8以前不支持
		DontSupportRenameIndex:    true, // 重命名索引 5.7 不支持
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(time.Second * 30)
	_db = db

	// 主从配置
	_ = _db.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{
			mysql.Open(connWrite),
		}, // 写操作
		Replicas: []gorm.Dialector{
			mysql.Open(connRead),
			mysql.Open(connWrite),
		}, // 读操作
		Policy: dbresolver.RandomPolicy{},
	}))

	Migration()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
