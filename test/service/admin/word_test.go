package admin_test

import (
	"goaway/pkg/library/util/business"
	service2 "goaway/pkg/service/admin"
	"testing"
)

func TestWordGetList(t *testing.T) {
	pager := business.NewPager()
	filter := &service2.Word{}
	filter.Word = "非人"
	list, total := service2.NewWord().WordGetList(pager, filter)
	t.Log(*pager)
	t.Log(total)
	for _, v := range list {
		t.Log(v)
	}
}

func TestWordGetListByIds(t *testing.T) {
	ids := []int{349047, 95763, 95615}
	list, err := service2.NewWord().WordGetListByIds(ids)
	t.Log(err)
	for _, v := range list {
		t.Log(v)
	}
}

func TestWordCreate(t *testing.T) {
	//service.Aop.SetCreateBefore(func(dataIn interface{}) {
	//	fmt.Println(dataIn)
	//})
	//service.Aop.SetCreateAfter(func(dataOut interface{}) {
	//	fmt.Println(dataOut)
	//})
	dataIn := service2.WordCreateForm{
		Word: "test",
	}
	dataOut, err := service2.NewWord().WordCreate(dataIn)
	if err != nil {
		return
	}
	t.Log(dataOut)
}

func TestWordGetInfoByIds(t *testing.T) {
	ids := make([]int, 0)
	ids = append(ids, 1)
	ids = append(ids, 2)
	ids = append(ids, 3)
	byIds, err := service2.NewWord().WordGetInfoByIds(ids)
	if err != nil {
		return
	}
	t.Log(byIds)
}

func TestWordUpdate(t *testing.T) {
	dataIn := service2.WordUpdateForm{
		Word:      "www",
		Frequency: 1,
	}
	dataOut := service2.NewWord().WordUpdateById(332149, dataIn)
	if dataOut {
		t.Log(dataOut)
	} else {
		t.Fatal(dataOut)
	}
}

func TestWordDelete(t *testing.T) {
	ids := make([]int, 0)
	ids = append(ids, 332149)
	dataIn := service2.WordMultipleDeleteForm{
		Ids: ids,
	}
	ret := service2.NewWord().WordDeleteByIds(dataIn)
	if ret {
		t.Log(ret)
	} else {
		t.Fatal(ret)
	}
}
