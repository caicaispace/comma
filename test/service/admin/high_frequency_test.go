package admin_test

import (
	service2 "comma/pkg/service/admin"
	"testing"

	"github.com/caicaispace/gohelper/business"
)

func TestHighFrequencyGetList(t *testing.T) {
	pager := business.NewPager()
	list, total := service2.NewHighFrequency().HighFrequencyGetList(pager)
	t.Log(*pager)
	t.Log(total)
	for _, v := range list {
		t.Log(v)
	}
}

func TestHighFrequencyCreate(t *testing.T) {
	dataIn := service2.HighFrequencyCreateForm{
		WordId: 1,
	}
	dataOut, err := service2.NewHighFrequency().HighFrequencyCreate(dataIn)
	if err != nil {
		return
	}
	t.Log(dataOut)
}

func TestHighFrequencyUpdate(t *testing.T) {
	dataIn := service2.HighFrequencyUpdateForm{
		WordId: 1,
	}
	dataOut := service2.NewHighFrequency().HighFrequencyUpdateById(352957, dataIn)
	if dataOut {
		t.Log(dataOut)
	} else {
		t.Fatal(dataOut)
	}
}
