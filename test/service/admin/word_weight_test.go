package admin_test

import (
	"comma/pkg/library/util/business"
	service2 "comma/pkg/service/admin"
	"testing"
)

func TestWordWeightGetList(t *testing.T) {
	filter := service2.Word{}
	pager := business.NewPager()
	list, total := service2.NewWordWeight().WordWeightGetList(pager, &filter)
	t.Log(*pager)
	t.Log(total)
	for _, v := range list {
		t.Log(v)
	}
}

func TestWordWeightCreate(t *testing.T) {
	dataIn := service2.WordWeightCreateForm{
		WordId: 1,
	}
	dataOut, err := service2.NewWordWeight().WordWeightCreate(dataIn)
	if err != nil {
		return
	}
	t.Log(dataOut)
}

func TestWordWeightUpdate(t *testing.T) {
	dataIn := service2.WordWeightUpdateForm{
		WordId: 1,
	}
	err := service2.NewWordWeight().WordWeightUpdateById(352957, dataIn)
	if err != nil {
		t.Log(err)
	}
}
