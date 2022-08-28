package collection

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	a1 := &Item{A: "a1"}
	a2 := &Item{A: "a2"}
	a3 := &Item{A: "a3"}
	f := NewFirst([]*Item{a1, a2})
	f.Add(a3)
	newColl := make([]*Item, 0)
	f.GetDataForPage(2, func(item any, key int) any {
		newColl = append(newColl, item.(*Item))
		return nil
	})
	fmt.Println(newColl[0].A)
}
