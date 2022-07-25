package admin_test

import (
	"testing"

	"github.com/caicaispace/gohelper/business"

	service2 "comma/pkg/service/admin"
)

func TestFestivalModelGetList(t *testing.T) {
	pager := business.NewPager()
	list, total := service2.NewFestival().FestivalGetList(pager)
	t.Log(*pager)
	t.Log(total)
	for _, v := range list {
		t.Log(v)
	}
}

func TestFestivalModelCreate(t *testing.T) {
	dataIn := service2.FestivalCreateForm{
		WordId: 1,
	}
	create, err := service2.NewFestival().FestivalCreate(dataIn)
	if err != nil {
		return
	}
	t.Log(create)
}

func TestFestivalModelUpdate(t *testing.T) {
	dataIn := service2.FestivalUpdateForm{
		WordId: 1,
	}
	outData, err := service2.NewFestival().FestivalUpdateById(352957, dataIn)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(outData)
}
