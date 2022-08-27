package collection

import (
	"github.com/jianfengye/collection"
)

type First struct {
	*Base
}

func New(objs any) *First {
	base := &Base{
		PageLimit:  2,
		PageTotal:  5,
		Collection: collection.NewObjCollection(objs),
	}
	base.MaxCount = base.PageLimit * base.PageTotal
	return &First{
		Base: base,
	}
}

func (f *First) Add(item *Item) bool {
	if f.Collection == nil {
		f.Collection = collection.NewObjCollection(item)
	}
	if f.Collection.Count() > f.MaxCount {
		return false
	}
	f.Collection.Append(item)
	return true
}

func (f *First) Handle() {
}
