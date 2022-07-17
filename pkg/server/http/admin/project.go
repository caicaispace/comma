package admin

import (
	"net/http"

	"comma/pkg/library/core/e"
	http2 "comma/pkg/library/net/http"
	"comma/pkg/service/admin"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func ProjectList(c *gin.Context) {
	ctx := http2.Context{C: c}
	pager := ctx.GetPager()
	list, _ := admin.NewProject().ProjectGetList(pager)
	ctx.Success(gin.H{
		"list":  list,
		"pager": pager.ToMap(),
	}, nil)
}

func ProjectCreate(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.ProjectCreateForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	outData, err := admin.NewProject().ProjectCreate(form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), nil)
		return
	}
	ctx.Success(gin.H{
		"info": outData,
	}, nil)
}

func ProjectUpdate(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.ProjectUpdateForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	outData, err := admin.NewProject().ProjectUpdateById(updateId, form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), err)
		return
	}
	ctx.Success(gin.H{
		"info": outData,
	}, nil)
}

func ProjectDelete(c *gin.Context) {
	ctx := http2.Context{C: c}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	ids := make([]int, 0)
	ids = append(ids, updateId)
	form := admin.ProjectMultipleDeleteForm{
		Ids: ids,
	}
	success := admin.NewProject().ProjectDeleteByIds(form)
	if !success {
		ctx.Error(http.StatusOK, e.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}

func ProjectMultipleDelete(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.ProjectMultipleDeleteForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	success := admin.NewProject().ProjectDeleteByIds(form)
	if !success {
		ctx.Error(http.StatusOK, e.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}
