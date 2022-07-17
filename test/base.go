package test

import (
	"comma/pkg/library/db"
	"comma/pkg/library/setting"
)

func init() {
	db.New(&setting.DBSetting{
		Host:     "127.0.0.1",
		Port:     "3305",
		Username: "root",
		Password: "123456",
		DbName:   "comma",
	})
}
