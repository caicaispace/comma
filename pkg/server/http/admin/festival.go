package admin

import (
	"comma/pkg/library/core/e"
	"comma/pkg/service/admin"
	"net/http"

	http2 "comma/pkg/library/net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func FestivalList(c *gin.Context) {
	ctx := http2.Context{C: c}
	pager := ctx.GetPager()
	list, _ := admin.NewFestival().FestivalGetList(pager)
	ctx.Success(gin.H{
		"list":  list,
		"pager": pager.ToMap(),
	}, nil)
}

func FestivalCreate(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.FestivalCreateForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	outData, err := admin.NewFestival().FestivalCreate(form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), nil)
		return
	}
	ctx.Success(gin.H{
		"info": outData,
	}, nil)
}

func FestivalUpdate(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.FestivalUpdateForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	outData, err := admin.NewFestival().FestivalUpdateById(updateId, form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), err)
		return
	}
	ctx.Success(gin.H{
		"info": outData,
	}, nil)
}

func FestivalDelete(c *gin.Context) {
	ctx := http2.Context{C: c}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	ids := make([]int, 0)
	ids = append(ids, updateId)
	form := admin.FestivalMultipleDeleteForm{
		Ids: ids,
	}
	err := admin.NewFestival().FestivalDeleteByIds(form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), err)
		return
	}
	ctx.Success(nil, nil)
}

func FestivalMultipleDelete(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.FestivalMultipleDeleteForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	err := admin.NewFestival().FestivalDeleteByIds(form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), err)
		return
	}
	ctx.Success(nil, nil)
}
