package admin

import (
	"net/http"

	"comma/pkg/service/admin"

	"github.com/caicaispace/gohelper/errx"

	httpServer "github.com/caicaispace/gohelper/server/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func HighFrequencyList(c *gin.Context) {
	ctx := httpServer.Context{C: c}
	pager := ctx.GetPager()
	list, _ := admin.NewHighFrequency().HighFrequencyGetList(pager)
	ctx.Success(gin.H{
		"list":  list,
		"pager": pager.ToMap(),
	}, nil)
}

func HighFrequencyCreate(c *gin.Context) {
	var (
		ctx  = httpServer.Context{C: c}
		form admin.HighFrequencyCreateForm
	)
	httpCode, errCode := httpServer.BindAndValid(c, &form)
	if errCode != errx.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	outData, err := admin.NewHighFrequency().HighFrequencyCreate(form)
	if err != nil {
		ctx.Error(http.StatusOK, errx.Error, err.Error(), nil)
		return
	}
	ctx.Success(gin.H{
		"info": outData,
	}, nil)
}

func HighFrequencyUpdate(c *gin.Context) {
	var (
		ctx  = httpServer.Context{C: c}
		form admin.HighFrequencyUpdateForm
	)
	httpCode, errCode := httpServer.BindAndValid(c, &form)
	if errCode != errx.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	success := admin.NewHighFrequency().HighFrequencyUpdateById(updateId, form)
	if !success {
		ctx.Error(http.StatusOK, errx.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}

func HighFrequencyDelete(c *gin.Context) {
	ctx := httpServer.Context{C: c}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	ids := make([]int, 0)
	ids = append(ids, updateId)
	form := admin.HighFrequencyMultipleDeleteForm{
		Ids: ids,
	}
	success := admin.NewHighFrequency().HighFrequencyDeleteByIds(form)
	if !success {
		ctx.Error(http.StatusOK, errx.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}

func HighFrequencyMultipleDelete(c *gin.Context) {
	var (
		ctx  = httpServer.Context{C: c}
		form admin.HighFrequencyMultipleDeleteForm
	)
	httpCode, errCode := httpServer.BindAndValid(c, &form)
	if errCode != errx.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	success := admin.NewHighFrequency().HighFrequencyDeleteByIds(form)
	if !success {
		ctx.Error(http.StatusOK, errx.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}
