package banned

import (
	"goaway/pkg/library/core/e"
	"goaway/pkg/library/core/t"
	http2 "goaway/pkg/library/net/http"
	service "goaway/pkg/service/banned"

	"github.com/gin-gonic/gin"
)

func Find(c *gin.Context) {
	var (
		ctx        = http2.Context{C: c}
		word       = c.Query("word")
		handleType = c.Query("type")
	)
	hasFind, textFindSlice := service.GetInstance().Find(word, handleType)
	rspData := t.Map{
		"has_find": hasFind,
		"text":     textFindSlice,
	}
	ctx.Success(rspData, nil)
}

type FindAddForm struct {
	Word string `form:"word" valid:"Required;MaxSize(32)"`
}

func Add(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form FindAddForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	service.GetInstance().Add(form.Word)
	ctx.Success(nil, nil)
}

type FindDelForm struct {
	Word string `form:"word" valid:"Required;MaxSize(32)"`
}

func Del(c *gin.Context) {
	var (
		ctx  = http2.Context{C: c}
		form FindDelForm
	)
	httpCode, errCode := http2.BindAndValid(c, &form)
	if errCode != e.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	service.GetInstance().Del(form.Word)
	ctx.Success(nil, nil)
}