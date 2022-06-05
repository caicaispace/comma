package admin_test

import (
	"goaway/pkg/library/db"
	"goaway/pkg/library/setting"
)

func init() {
	db.New(&setting.DBSetting{
		Username: "root",
		Password: "123456",
		Host:     "127.0.0.1",
		DbName:   "goaway",
	})
}
