package admin

import (
	"net/http"

	"comma/pkg/library/core/e"
	http2 "comma/pkg/library/net/http"
	"comma/pkg/service/admin"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func HyponymList(c *gin.Context) {
	ctx := http2.Context{C: c}
	pager := ctx.GetPager()
	list, _ := admin.NewHyponym().HyponymGetList(pager)
	ctx.Success(gin.H{
		"list":  list,
		"pager": pager.ToMap(),
	}, nil)
}

func HyponymCreate(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.HyponymCreateForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	outData, err := admin.NewHyponym().HyponymCreate(form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), nil)
		return
	}
	ctx.Success(gin.H{
		"info": outData,
	}, nil)
}

func HyponymUpdate(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.HyponymUpdateForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	success := admin.NewHyponym().HyponymUpdateById(updateId, form)
	if !success {
		ctx.Error(http.StatusOK, e.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}

func HyponymDelete(c *gin.Context) {
	ctx := http2.Context{C: c}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	ids := make([]int, 0)
	ids = append(ids, updateId)
	form := admin.HyponymMultipleDeleteForm{
		Ids: ids,
	}
	success := admin.NewHyponym().HyponymDeleteByIds(form)
	if !success {
		ctx.Error(http.StatusOK, e.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}

func HyponymMultipleDelete(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.HyponymMultipleDeleteForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	success := admin.NewHyponym().HyponymDeleteByIds(form)
	if !success {
		ctx.Error(http.StatusOK, e.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}
