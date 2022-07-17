package admin_test

import (
	"comma/pkg/library/util/business"
	service2 "comma/pkg/service/admin"
	"testing"
)

func TestSynonymGetList(t *testing.T) {
	pager := business.NewPager()
	list, total := service2.NewSynonym().SynonymGetList(pager)
	t.Log(*pager)
	t.Log(total)
	for _, v := range list {
		t.Log(v)
	}
}

func TestSynonymCreate(t *testing.T) {
	dataIn := service2.SynonymCreateForm{
		WordIds: "1111,2222",
	}
	dataOut, err := service2.NewSynonym().SynonymCreate(dataIn)
	if err != nil {
		return
	}
	t.Log(dataOut)
}

func TestSynonymUpdate(t *testing.T) {
	dataIn := service2.SynonymUpdateForm{
		WordIds: "1111,2222",
	}
	outData, err := service2.NewSynonym().SynonymUpdateById(352957, dataIn)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(outData)
}
