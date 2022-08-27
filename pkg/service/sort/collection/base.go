package collection

import (
	"github.com/caicaispace/gohelper/mathx"
	"github.com/jianfengye/collection"
)

type Item struct {
	A string
}

type Base struct {
	MaxCount   int
	PageLimit  int
	PageTotal  int
	Collection *collection.ObjCollection
}

func (b *Base) IsFull() bool {
	return b.Collection.Count() > b.MaxCount
}

func (b *Base) GetDataForPage(page int, callback func(item any, key int) any) {
	// offset := mathx.MaxInt(0, (page-1)*b.PageLimit)
	page = mathx.MaxInt(0, page-1)
	b.Collection.ForPage(page, b.PageLimit).Map(callback)
}
