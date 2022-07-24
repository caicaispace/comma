package admin_test

import (
	"testing"

	"comma/pkg/library/util/business"
	service2 "comma/pkg/service/admin"
)

func TestBannedGetList(t *testing.T) {
	filter := service2.Word{}
	pager := business.NewPager()
	list, total := service2.NewBanned().BannedGetList(pager, &filter)
	t.Log(*pager)
	t.Log(total)
	for _, v := range list {
		t.Log(v)
	}
}

func TestBannedCreate(t *testing.T) {
	form := service2.BannedCreateForm{
		WordId:    779,
		ProjectId: 1,
	}
	create, err := service2.NewBanned().BannedCreate(form)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(create)
}

func TestBannedUpdate(t *testing.T) {
	form := service2.BannedUpdateForm{
		// WordId: 1,
	}
	err := service2.NewBanned().BannedUpdateById(352957, form)
	if err != nil {
		t.Log(err)
	} else {
		t.Fatal(err)
	}
}
