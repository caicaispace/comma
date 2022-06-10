package db

import (
	"fmt"
	"os"

	"goaway/pkg/library/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db *gorm.DB

// New https://github.com/moell-peng/gin-gorm-example
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
	// slowLogger := logger.New(
	// 	// stdout
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
	// 	logger.Config{
	// 		SlowThreshold: time.Second,   // 慢 SQL 阈值
	// 		LogLevel:      logger.Silent, // Log level
	// 		Colorful:      false,         // 禁用彩色打印
	// 	},
	// )
	return &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// Logger: slowLogger,
	}
}
