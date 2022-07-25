package db

import (
	"comma/pkg/library/setting"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db *gorm.DB

func New(conf *setting.DBSetting) {
	NewWithAddr(conf.ToAddrString())
}

func NewWithAddr(addr string) {
	db, err := gorm.Open(mysql.Open(addr), getGormConfig())
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	_db = db
}

func DB() *gorm.DB {
	return _db
}

func getGormConfig() *gorm.Config {
	c := &gorm.Config{}
	if setting.Server.Env == "dev" {
		c.Logger = logger.Default.LogMode(logger.Info)
		// c.Logger slowLogger()
	}
	return c
}

func slowLogger() {
	// return logger.New(
	// 	// stdout
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
	// 	logger.Config{
	// 		SlowThreshold: time.Second,   // 慢 SQL 阈值
	// 		LogLevel:      logger.Silent, // Log level
	// 		Colorful:      false,         // 禁用彩色打印
	// 	},
	// )
}
