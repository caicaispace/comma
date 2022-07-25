package banned

import (
	"github.com/caicaispace/gohelper/errx"
	"github.com/caicaispace/gohelper/syntax"

	httpServer "github.com/caicaispace/gohelper/server/http"

	service "comma/pkg/service/banned"

	"github.com/gin-gonic/gin"
)

func Find(c *gin.Context) {
	var (
		ctx        = httpServer.Context{C: c}
		word       = c.Query("word")
		handleType = c.Query("type")
	)
	hasFind, textFindSlice := service.GetInstance().Find(word, handleType)
	rspData := syntax.Map{
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
		ctx  = httpServer.Context{C: c}
		form FindAddForm
	)
	httpCode, errCode := httpServer.BindAndValid(c, &form)
	if errCode != errx.Success {
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
		ctx  = httpServer.Context{C: c}
		form FindDelForm
	)
	httpCode, errCode := httpServer.BindAndValid(c, &form)
	if errCode != errx.Success {
		ctx.Error(httpCode, errCode, nil, nil)
		return
	}
	service.GetInstance().Del(form.Word)
	ctx.Success(nil, nil)
}
