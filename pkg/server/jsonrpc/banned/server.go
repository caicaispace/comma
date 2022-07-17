package banned

import (
	"sync"

	service "comma/pkg/service/banned"
)

type Banned struct{}

var (
	s    *Banned
	once sync.Once
)

func GetInstance() *Banned {
	once.Do(func() {
		s = &Banned{}
	})
	return s
}

type findParams struct {
	Word string `json:"word"`
	Type string `json:"type"`
}

type findResult struct {
	HasFind bool     `json:"has_find"`
	Text    []string `json:"text"`
}

func (*Banned) Find(params *findParams, result *findResult) error {
	hasFind, textFindSlice := service.GetInstance().Find(params.Word, params.Type)
	rspData := findResult{
		HasFind: hasFind,
		Text:    textFindSlice,
	}
	*result = interface{}(rspData).(findResult)
	return nil
}

type addParams struct {
	Word string `json:"word"`
}

func (*Banned) Add(params *addParams, result *interface{}) error {
	service.GetInstance().Add(params.Word)
	return nil
}

type delParams struct {
	Word string `json:"word"`
}

func (*Banned) Del(params *addParams, result *interface{}) error {
	service.GetInstance().Del(params.Word)
	return nil
}
