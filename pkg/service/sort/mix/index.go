package mix

import (
	"fmt"

	"comma/pkg/service/sort/mix/collection"
)

type MixSort struct {
	items []*collection.Item
	first *collection.First
}

func New() *MixSort {
	return &MixSort{
		first: &collection.First{},
	}
}

func (ms *MixSort) AddItems(items []*collection.Item) *MixSort {
	ms.items = items
	for _, item := range items {
		ms.first.Add(item)
	}
	return ms
}

func (ms *MixSort) GetDataForPage(page int) []*collection.Item {
	if page > 5 {
		return nil
	}
	list := make([]*collection.Item, 0)
	ms.first.GetDataForPage(page, func(item any, key int) any {
		list = append(list, item.(*collection.Item))
		return nil
	})
	fmt.Println(list[0].A)
	return list
}
