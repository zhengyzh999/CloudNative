package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB
var datasourceUrl = "root:Down@2022@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               datasourceUrl,
		DefaultStringSize: 256,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}
	setPool(DB)
}

func setPool(DB *gorm.DB) {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	// 最大连接数
	sqlDB.SetMaxOpenConns(10)
	// 最大空闲连接数
	sqlDB.SetMaxIdleConns(5)
	// 连接最大活跃时间
	sqlDB.SetConnMaxLifetime(time.Hour)
}
