package controller

import (
	"comma/pkg/library/util/business"
)

func Pager() *business.Pager {
	return business.NewPager()
}
