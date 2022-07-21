package admin

import (
	"comma/pkg/library/core/e"
	"comma/pkg/service/admin"
	"net/http"

	http2 "comma/pkg/library/net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func WeightList(c *gin.Context) {
	ctx := http2.Context{C: c}
	pager := ctx.GetPager()
	filter := &admin.Word{}
	if ctx.C.Query("word") != "" {
		filter.Word = ctx.C.Query("word")
	}
	list, _ := admin.NewWordWeight().WordWeightGetList(pager, filter)
	ctx.Success(gin.H{
		"list":  list,
		"pager": pager.ToMap(),
	}, nil)
}

func WeightCreate(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.WordWeightCreateForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	outData, err := admin.NewWordWeight().WordWeightCreate(form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), nil)
		return
	}
	ctx.Success(gin.H{
		"info": outData,
	}, nil)
}

func WeightUpdate(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.WordWeightUpdateForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	err := admin.NewWordWeight().WordWeightUpdateById(updateId, form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), err)
		return
	}
	ctx.Success(nil, nil)
}

func WeightDelete(c *gin.Context) {
	ctx := http2.Context{C: c}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	ids := make([]int, 0)
	ids = append(ids, updateId)
	form := admin.WordWeightMultipleDeleteForm{
		Ids: ids,
	}
	err := admin.NewWordWeight().WordWeightDeleteByIds(form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), err)
		return
	}
	ctx.Success(nil, nil)
}

func WeightMultipleDelete(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.WordWeightMultipleDeleteForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	err := admin.NewWordWeight().WordWeightDeleteByIds(form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), err)
		return
	}
	ctx.Success(nil, nil)
}
