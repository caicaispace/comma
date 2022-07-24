package admin

import (
	"net/http"

	"comma/pkg/library/core/e"
	"comma/pkg/service/admin"

	http2 "comma/pkg/library/net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func BannedList(c *gin.Context) {
	ctx := http2.Context{C: c}
	pager := ctx.GetPager()
	filter := &admin.Word{}
	if ctx.C.Query("word") != "" {
		filter.Word = ctx.C.Query("word")
	}
	list, _ := admin.NewBanned().BannedGetList(pager, filter)
	ctx.Success(gin.H{
		"list":  list,
		"pager": pager.ToMap(),
	}, nil)
}

func BannedCreate(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.BannedCreateForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	outData, err := admin.NewBanned().BannedCreate(form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), nil)
		return
	}
	ctx.Success(gin.H{
		"info": outData,
	}, nil)
}

func BannedUpdate(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.BannedUpdateForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	err := admin.NewBanned().BannedUpdateById(updateId, form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), err)
		return
	}
	ctx.Success(nil, nil)
}

func BannedDelete(c *gin.Context) {
	ctx := http2.Context{C: c}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	ids := make([]int, 0)
	ids = append(ids, updateId)
	form := admin.BannedMultipleDeleteForm{
		Ids: ids,
	}
	err := admin.NewBanned().BannedDeleteByIds(form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), err)
		return
	}
	ctx.Success(nil, nil)
}

func BannedMultipleDelete(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.BannedMultipleDeleteForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	err := admin.NewBanned().BannedDeleteByIds(form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), err)
		return
	}
	ctx.Success(nil, nil)
}
