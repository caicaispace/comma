package controller

import (
	"goaway/pkg/library/util/business"
)

func Pager() *business.Pager {
	return business.NewPager()
}
