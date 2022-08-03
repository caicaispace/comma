package admin

import (
	"comma/pkg/service/admin"
	"net/http"

	"github.com/caicaispace/gohelper/errx"

	httpServer "github.com/caicaispace/gohelper/server/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func ProjectList(c *gin.Context) {
	ctx := httpServer.Context{C: c}
	pager := ctx.GetPager()
	list, _ := admin.NewProject().ProjectGetList(pager)
	ctx.Success(gin.H{
		"list":  list,
		"pager": pager.ToMap(),
	}, nil)
}

func ProjectCreate(c *gin.Context) {
	var (
		ctx  = httpServer.Context{C: c}
		form admin.ProjectCreateForm
	)
	httpCode, errCode := httpServer.BindAndValid(c, &form)
	if errCode != errx.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	outData, err := admin.NewProject().ProjectCreate(form)
	if err != nil {
		ctx.Error(http.StatusOK, errx.Error, err.Error(), nil)
		return
	}
	ctx.Success(gin.H{
		"info": outData,
	}, nil)
}

func ProjectUpdate(c *gin.Context) {
	var (
		ctx  = httpServer.Context{C: c}
		form admin.ProjectUpdateForm
	)
	httpCode, errCode := httpServer.BindAndValid(c, &form)
	if errCode != errx.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	outData, err := admin.NewProject().ProjectUpdateById(updateId, form)
	if err != nil {
		ctx.Error(http.StatusOK, errx.Error, err.Error(), err)
		return
	}
	ctx.Success(gin.H{
		"info": outData,
	}, nil)
}

func ProjectDelete(c *gin.Context) {
	ctx := httpServer.Context{C: c}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	ids := make([]int, 0)
	ids = append(ids, updateId)
	form := admin.ProjectMultipleDeleteForm{
		Ids: ids,
	}
	success := admin.NewProject().ProjectDeleteByIds(form)
	if !success {
		ctx.Error(http.StatusOK, errx.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}

func ProjectMultipleDelete(c *gin.Context) {
	var (
		ctx  = httpServer.Context{C: c}
		form admin.ProjectMultipleDeleteForm
	)
	httpCode, errCode := httpServer.BindAndValid(c, &form)
	if errCode != errx.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	success := admin.NewProject().ProjectDeleteByIds(form)
	if !success {
		ctx.Error(http.StatusOK, errx.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}
