package controller

import (
	"fmt"

	"comma/pkg/library/db"

	"github.com/caicaispace/gohelper/server/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Test(c *gin.Context) {
	ctx := http.Context{C: c}
	ctx.Success(nil, nil)
}

func TestPager(c *gin.Context) {
	ctx := http.Context{C: c}
	pager := ctx.GetPager()
	pager.SetTotal(100)
	ctx.Success(gin.H{
		"page":  pager.GetPage(),
		"limit": pager.GetLimit(),
		"total": pager.GetTotal(),
	}, nil)
}

type Menu struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func DB(c *gin.Context) {
	ctx := http.Context{C: c}
	results := make([]*Menu, 0)
	db.DB().Table("menu").Where("id > ?", 0).FindInBatches(&results, 5, func(tx *gorm.DB, batch int) error {
		for _, result := range results {
			fmt.Println(result)
		}
		return nil
	})
	ctx.Success(nil, nil)
}
