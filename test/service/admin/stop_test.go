package admin_test

import (
	"comma/pkg/library/util/business"
	service2 "comma/pkg/service/admin"
	"testing"
)

func TestStopGetList(t *testing.T) {
	filter := service2.Word{}
	pager := business.NewPager()
	list, total := service2.NewStop().StopGetList(pager, &filter)
	t.Log(*pager)
	t.Log(total)
	for _, v := range list {
		t.Log(v)
	}
}

func TestStopCreate(t *testing.T) {
	dataIn := service2.StopCreateForm{
		WordId:    1,
		ProjectId: 1,
	}
	dataOut, err := service2.NewStop().StopCreate(dataIn)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(dataOut)
}

func TestStopUpdate(t *testing.T) {
	dataIn := service2.StopUpdateForm{
		WordId:    1,
		ProjectId: 1,
	}
	outData, err := service2.NewStop().StopUpdateById(352957, dataIn)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(outData)
}
