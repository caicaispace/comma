package business

import (
	"sync"
)

type Pager struct {
	page  int
	limit int
	total int
}

const (
	defaultPage  = 1
	defaultLimit = 10
	defaultTotal = 0
)

var (
	p    *Pager
	once sync.Once
)

func GetInstance() *Pager {
	once.Do(func() {
		p = NewPager()
	})
	return p
}

func NewPager() *Pager {
	return &Pager{
		page:  defaultPage,
		limit: defaultLimit,
		total: defaultTotal,
	}
}

func (p *Pager) ToMap() *map[string]int {
	return &map[string]int{
		"p_page":  p.page,
		"p_limit": p.limit,
		"p_total": p.total,
	}
}

func (p *Pager) SetPage(page int) *Pager {
	p.page = page
	return p
}

func (p *Pager) SetLimit(limit int) *Pager {
	p.limit = limit
	return p
}

func (p *Pager) SetTotal(total int) *Pager {
	p.total = total
	return p
}

func (p *Pager) GetPage() int {
	if p.page == 0 {
		return defaultPage
	}
	return p.page
}

func (p *Pager) GetLimit() int {
	if p.limit == 0 {
		p.limit = defaultLimit
	}
	return p.limit
}

func (p *Pager) GetOffset() int {
	if p.page == 1 {
		return 0
	}
	return (p.page - 1) * p.limit
}

func (p *Pager) GetTotal() int {
	return p.total
}
