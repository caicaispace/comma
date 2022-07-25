package pinyin

import (
	"comma/pkg/service/admin"
	"fmt"
	"strings"
	"sync"

	"github.com/mozillazg/go-pinyin"
)

type pinyinService struct{}

var (
	service *pinyinService
	once    sync.Once
)

func GetInstance() *pinyinService {
	once.Do(func() {
		service = &pinyinService{}
	})
	return service
}

type ToPinyinOut struct {
	Keyword  string
	Pinyin   string
	Initials string
	Pinyins  []string
}

func (ps *pinyinService) ToPinyin(word string) *ToPinyinOut {
	outData := &ToPinyinOut{}
	words := pinyin.LazyPinyin(word, pinyin.NewArgs())
	outData.Keyword = word
	outData.Pinyins = words
	outData.Pinyin = strings.Join(words, "")
	outData.Initials = ""
	for _, word := range words {
		outData.Initials += word[0:1]
	}
	err := ps.Create(outData)
	if err != nil {
		fmt.Println(err)
	}
	return outData
}

func (ps *pinyinService) Create(inData *ToPinyinOut) error {
	outData := admin.PinyinCreateForm{
		Keyword:        inData.Keyword,
		Pinyin:         inData.Pinyin,
		Pinyins:        strings.Join(inData.Pinyins, ","),
		PinyinInitials: inData.Initials,
		IsDel:          0,
	}
	curd := admin.NewPinyin()
	_, err := curd.PinyinCreate(outData)
	if err != nil {
		return err
	}
	return nil
}
