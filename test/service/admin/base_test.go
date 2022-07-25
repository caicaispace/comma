package admin_test

import (
	"comma/pkg/library/db"

	"github.com/caicaispace/gohelper/setting"
)

func init() {
	db.New(&setting.DBSetting{
		Username: "root",
		Password: "123456",
		Host:     "127.0.0.1",
		DbName:   "comma",
	})
}
