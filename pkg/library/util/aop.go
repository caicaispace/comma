package util

const (
	CreateBefore = "createBefore"
	CreateAfter  = "createAfter"
)

type Aop struct {
	AopMethod map[string]func(interface{})
}

func (a *Aop) SetCreateBefore(fn func(interface{})) {
	a.AopMethod[CreateBefore] = fn
}

func (a *Aop) SetCreateAfter(fn func(interface{})) {
	a.AopMethod[CreateAfter] = fn
}

func (a *Aop) RunCreateBefore(inData interface{}) {
	if fn, ok := a.AopMethod[CreateBefore]; ok {
		fn(inData)
	}
}

func (a *Aop) RunCreateAfter(outData interface{}) {
	if fn, ok := a.AopMethod[CreateAfter]; ok {
		fn(outData)
	}
}
