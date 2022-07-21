package http

import (
	"comma/pkg/library/core/e"
	"comma/pkg/library/core/l"
	"comma/pkg/library/util/business"
	"net/http"

	"github.com/unknwon/com"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		l.Info(err.Key, err.Message)
	}
}

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, body interface{}) (int, int) {
	err := c.Bind(body)
	if err != nil {
		return http.StatusBadRequest, e.InvalidParams
	}
	valid := validation.Validation{}
	check, err := valid.Valid(body)
	if err != nil {
		return http.StatusInternalServerError, e.Error
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.InvalidParams
	}
	return http.StatusOK, e.Success
}

func (c *Context) GetPager() *business.Pager {
	pager := business.GetInstance()
	pager.SetPage(com.StrTo(c.C.Query("p_page")).MustInt())
	pager.SetLimit(com.StrTo(c.C.Query("p_limit")).MustInt())
	pager.SetTotal(com.StrTo(c.C.Query("p_total")).MustInt())
	return pager
}
