package http

import (
	"comma/pkg/library/core/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context struct {
	C *gin.Context
}

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ResponseData response setting gin.JSON
func (c *Context) ResponseData(httpCode, errCode int, data interface{}, msg interface{}, err error) {
	rsp := response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	}
	if msg != nil {
		rsp.Msg = msg.(string)
	}
	c.C.JSON(httpCode, rsp)
	if err != nil {
		_ = c.C.Error(err)
	}
}

func (c *Context) Success(data interface{}, err error) {
	c.ResponseData(http.StatusOK, e.Success, data, nil, err)
}

func (c *Context) Error(httpCode, errCode int, msg interface{}, err error) {
	c.ResponseData(httpCode, errCode, nil, msg, err)
}

func (c *Context) JOSN(body []byte) {
	c.C.Writer.Header().Set("Content-Type", "application/json")
	c.C.Writer.Write(body)
}
