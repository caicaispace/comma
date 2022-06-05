package admin_test

import (
	"goaway/pkg/library/util/business"
	service2 "goaway/pkg/service/admin"
	"testing"
)

func TestProjectGetList(t *testing.T) {
	pager := business.NewPager()
	list, total := service2.NewProject().ProjectGetList(pager)
	t.Log(*pager)
	t.Log(total)
	for _, v := range list {
		t.Log(v)
	}
}

func TestProjectCreate(t *testing.T) {
	dataIn := service2.ProjectCreateForm{
		Name: "test",
	}
	create, err := service2.NewProject().ProjectCreate(dataIn)
	if err != nil {
		return
	}
	t.Log(create)
}

func TestProjectUpdate(t *testing.T) {
	dataIn := service2.ProjectUpdateForm{
		Name: "test",
	}
	outData, err := service2.NewProject().ProjectUpdateById(352957, dataIn)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(outData)
}
