package db_test

import (
	"comma/pkg/library/db"
	"comma/pkg/library/setting"
	"comma/pkg/model"
	"fmt"
	"testing"

	"gorm.io/gorm"
)

type Menu struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func TestDB(t *testing.T) {
	config := &setting.DBSetting{
		Username: "root",
		Password: "123456",
		Host:     "127.0.0.1",
		DbName:   "comma",
	}
	db.New(config)

	results := make([]*Menu, 0)
	ret := db.DB().Table("menu").Where("id > ?", 0).FindInBatches(&results, 5, func(tx *gorm.DB, batch int) error {
		for _, result := range results {
			fmt.Println(result)
			// batch processing found records
		}
		// tx.Save(&results)
		// tx.RowsAffected // number of records in this batch
		// batch // Batch 1, 2, 3
		// returns error will stop future batches
		return nil
	})
	t.Log(ret)
	// DB().Raw("select * from menu where id > ?", 0).Scan(&results)
	// for _, result := range results {
	// 	fmt.Println(result)
	// }
}

func Test_banner(t *testing.T) {
	config := &setting.DBSetting{
		Username: "root",
		Password: "123456",
		Host:     "127.0.0.1",
		DbName:   "comma",
	}
	db.New(config)
	bannerMgr := model.DictBannedMgr(db.DB())
	list, _ := bannerMgr.Gets()
	for _, item := range list {
		fmt.Println(item)
	}
}
