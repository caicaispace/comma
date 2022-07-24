package admin

import (
	"net/http"

	"comma/pkg/library/core/e"
	"comma/pkg/service/admin"

	http2 "comma/pkg/library/net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func StopList(c *gin.Context) {
	ctx := http2.Context{C: c}
	filter := &admin.Word{}
	if ctx.C.Query("word") != "" {
		filter.Word = ctx.C.Query("word")
	}
	pager := ctx.GetPager()
	list, _ := admin.NewStop().StopGetList(pager, filter)
	ctx.Success(gin.H{
		"list":  list,
		"pager": pager.ToMap(),
	}, nil)
}

func StopCreate(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.StopCreateForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	outData, err := admin.NewStop().StopCreate(form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), nil)
		return
	}
	ctx.Success(gin.H{
		"info": outData,
	}, nil)
}

func StopUpdate(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.StopUpdateForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	outData, err := admin.NewStop().StopUpdateById(updateId, form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), err)
		return
	}
	ctx.Success(gin.H{
		"info": outData,
	}, nil)
}

func StopDelete(c *gin.Context) {
	ctx := http2.Context{C: c}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	ids := make([]int, 0)
	ids = append(ids, updateId)
	form := admin.StopMultipleDeleteForm{
		Ids: ids,
	}
	success := admin.NewStop().StopDeleteByIds(form)
	if !success {
		ctx.Error(http.StatusOK, e.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}

func StopMultipleDelete(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.StopMultipleDeleteForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	success := admin.NewStop().StopDeleteByIds(form)
	if !success {
		ctx.Error(http.StatusOK, e.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}
