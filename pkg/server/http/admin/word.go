package admin

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"goaway/pkg/library/core/e"
	http2 "goaway/pkg/library/net/http"
	"goaway/pkg/service/admin"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func WordList(c *gin.Context) {
	ctx := http2.Context{C: c}
	pager := ctx.GetPager()
	filter := &admin.Word{}
	if ctx.C.Query("word") != "" {
		filter.Word = ctx.C.Query("word")
	}
	list, _ := admin.NewWord().WordGetList(pager, filter)
	ctx.Success(gin.H{
		"list":  list,
		"pager": pager.ToMap(),
	}, nil)
}

func WordListByIds(c *gin.Context) {
	ctx := http2.Context{C: c}
	if ctx.C.Param("ids") == "" {
		ctx.Error(http.StatusOK, e.Error, nil, nil)
		return
	}
	idsStr := ctx.C.Param("ids")
	fmt.Println(idsStr)
	idsSlice := strings.Split(idsStr, ",")
	ids := make([]int, 0)
	for _, v := range idsSlice {
		id, _ := strconv.Atoi(v)
		ids = append(ids, id)
	}
	list, _ := admin.NewWord().WordGetListByIds(ids)
	ctx.Success(gin.H{
		"list": list,
	}, nil)
}

func WordInfo(c *gin.Context) {
	ctx := http2.Context{C: c}
	id := com.StrTo(ctx.C.Param("id")).MustInt()
	info, _ := admin.NewWord().WordGetInfoById(id)
	ctx.Success(gin.H{
		"info": info,
	}, nil)
}

func WordCreate(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.WordCreateForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	outData, err := admin.NewWord().WordCreate(form)
	if err != nil {
		ctx.Error(http.StatusOK, e.Error, err.Error(), nil)
		return
	}
	ctx.Success(gin.H{
		"info": outData,
	}, nil)
}

func WordUpdate(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.WordUpdateForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	success := admin.NewWord().WordUpdateById(updateId, form)
	if !success {
		ctx.Error(http.StatusOK, e.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}

func WordDelete(c *gin.Context) {
	ctx := http2.Context{C: c}
	updateId := com.StrTo(ctx.C.Param("id")).MustInt()
	ids := make([]int, 0)
	ids = append(ids, updateId)
	form := admin.WordMultipleDeleteForm{
		Ids: ids,
	}
	success := admin.NewWord().WordDeleteByIds(form)
	if !success {
		ctx.Error(http.StatusOK, e.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}

func WordMultipleDelete(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form admin.WordMultipleDeleteForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	success := admin.NewWord().WordDeleteByIds(form)
	if !success {
		ctx.Error(http.StatusOK, e.Error, nil, nil)
		return
	}
	ctx.Success(nil, nil)
}
