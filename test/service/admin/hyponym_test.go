package admin_test

import (
	"goaway/pkg/library/util/business"
	service2 "goaway/pkg/service/admin"
	"testing"
)

func TestHyponymGetList(t *testing.T) {
	pager := business.NewPager()
	list, total := service2.NewHyponym().HyponymGetList(pager)
	t.Log(*pager)
	t.Log(total)
	for _, v := range list {
		t.Log(v)
	}
}

func TestHyponymCreate(t *testing.T) {
	dataIn := service2.HyponymCreateForm{
		HypernymWordId: 1,
	}
	create, err := service2.NewHyponym().HyponymCreate(dataIn)
	if err != nil {
		return
	}
	t.Log(create)
}

func TestHyponymUpdate(t *testing.T) {
	dataIn := service2.HyponymUpdateForm{
		HypernymWordId: 1,
	}
	ret := service2.NewHyponym().HyponymUpdateById(352957, dataIn)
	if ret {
		t.Log(ret)
	} else {
		t.Fatal(ret)
	}
}
