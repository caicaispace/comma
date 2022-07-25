package controller

import (
	"github.com/caicaispace/gohelper/business"
)

func Pager() *business.Pager {
	return business.NewPager()
}
