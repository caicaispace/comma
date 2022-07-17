package admin

import (
	"net/http"

	"comma/pkg/library/core/e"
	http2 "comma/pkg/library/net/http"
	"comma/pkg/service/admin"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func HighFrequencyList(c *gin.Context) {
	ctx := http2.Context{C: c}
	pager := ctx.GetPager()
	list, _ := admin.NewHighFrequency().HighFrequencyGetList(pager)
	ctx.Success(gin.H{
		"list":  list,
		"pager": pager.ToMap(),
	}, nil)
}

func HighFrequencyCreate(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.HighFrequencyCreateForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	outData, err := admin.NewHighFrequency().HighFrequencyCreate(form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), nil)
		return
	}
	ctx.Success(gin.H{
		"info": outData,
	}, nil)
}

func HighFrequencyUpdate(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.HighFrequencyUpdateForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	success := admin.NewHighFrequency().HighFrequencyUpdateById(updateId, form)
	if !success {
		ctx.Error(http.StatusOK, e.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}

func HighFrequencyDelete(c *gin.Context) {
	ctx := http2.Context{C: c}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	ids := make([]int, 0)
	ids = append(ids, updateId)
	form := admin.HighFrequencyMultipleDeleteForm{
		Ids: ids,
	}
	success := admin.NewHighFrequency().HighFrequencyDeleteByIds(form)
	if !success {
		ctx.Error(http.StatusOK, e.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}

func HighFrequencyMultipleDelete(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.HighFrequencyMultipleDeleteForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	success := admin.NewHighFrequency().HighFrequencyDeleteByIds(form)
	if !success {
		ctx.Error(http.StatusOK, e.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}
