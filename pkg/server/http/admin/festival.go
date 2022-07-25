package admin

import (
	"net/http"

	"comma/pkg/service/admin"

	"github.com/caicaispace/gohelper/errx"

	httpServer "github.com/caicaispace/gohelper/server/http"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func FestivalList(c *gin.Context) {
	ctx := httpServer.Context{C: c}
	pager := ctx.GetPager()
	list, _ := admin.NewFestival().FestivalGetList(pager)
	ctx.Success(gin.H{
		"list":  list,
		"pager": pager.ToMap(),
	}, nil)
}

func FestivalCreate(c *gin.Context) {
	var (
		ctx  = httpServer.Context{C: c}
		form admin.FestivalCreateForm
	)
	httpCode, errCode := httpServer.BindAndValid(c, &form)
	if errCode != errx.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	outData, err := admin.NewFestival().FestivalCreate(form)
	if err != nil {
		ctx.Error(http.StatusOK, errx.Error, err.Error(), nil)
		return
	}
	ctx.Success(gin.H{
		"info": outData,
	}, nil)
}

func FestivalUpdate(c *gin.Context) {
	var (
		ctx  = httpServer.Context{C: c}
		form admin.FestivalUpdateForm
	)
	httpCode, errCode := httpServer.BindAndValid(c, &form)
	if errCode != errx.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	outData, err := admin.NewFestival().FestivalUpdateById(updateId, form)
	if err != nil {
		ctx.Error(http.StatusOK, errx.Error, err.Error(), err)
		return
	}
	ctx.Success(gin.H{
		"info": outData,
	}, nil)
}

func FestivalDelete(c *gin.Context) {
	ctx := httpServer.Context{C: c}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	ids := make([]int, 0)
	ids = append(ids, updateId)
	form := admin.FestivalMultipleDeleteForm{
		Ids: ids,
	}
	err := admin.NewFestival().FestivalDeleteByIds(form)
	if err != nil {
		ctx.Error(http.StatusOK, errx.Error, err.Error(), err)
		return
	}
	ctx.Success(nil, nil)
}

func FestivalMultipleDelete(c *gin.Context) {
	var (
		ctx  = httpServer.Context{C: c}
		form admin.FestivalMultipleDeleteForm
	)
	httpCode, errCode := httpServer.BindAndValid(c, &form)
	if errCode != errx.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	err := admin.NewFestival().FestivalDeleteByIds(form)
	if err != nil {
		ctx.Error(http.StatusOK, errx.Error, err.Error(), err)
		return
	}
	ctx.Success(nil, nil)
}
