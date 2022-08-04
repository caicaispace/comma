package admin_test

import (
	"github.com/caicaispace/gohelper/orm/gorm"
	"github.com/caicaispace/gohelper/setting"
)

func init() {
	config := &setting.DBSetting{
		Username: "root",
		Password: "123456",
		Host:     "127.0.0.1",
		DbName:   "comma",
	}
	gorm.GetInstance().AddConnWithConfig(config, "")
}
