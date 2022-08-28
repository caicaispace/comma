package pinyin

import (
	"fmt"
	"sync"

	"comma/pkg/service/admin"

	"github.com/caicaispace/gohelper/tree/triedoublearray"
)

type Pinyin struct {
	pinyinTree           *triedoublearray.Cedar
	pinyinValues         [][]string
	pinyinInitialsTree   *triedoublearray.Cedar
	pinyinInitialsValues [][]string
}

var (
	service *Pinyin
	once    sync.Once
)

func GetInstance() *Pinyin {
	once.Do(func() {
		service = New()
	})
	return service
}

// Pinyin
func New() *Pinyin {
	py := &Pinyin{
		pinyinTree:           triedoublearray.New(),
		pinyinValues:         make([][]string, 0),
		pinyinInitialsTree:   triedoublearray.New(),
		pinyinInitialsValues: make([][]string, 0),
	}
	return py
}

func (py *Pinyin) FindByPinyinPrefix(pinyinInitials string) (map[string][]string, error) {
	ret := make(map[string][]string)
	for _, id := range py.pinyinInitialsTree.PrefixPredict([]byte(pinyinInitials), 0) {
		key, err := py.pinyinInitialsTree.Key(id)
		if err != nil {
			return nil, err
		}
		value, err := py.pinyinInitialsTree.Value(id)
		if err != nil {
			return nil, err
		}
		ret[string(key)] = py.pinyinInitialsValues[value]
	}
	return ret, nil
}

func (py *Pinyin) FindByPinyin(pinyin string) (map[string][]string, error) {
	ret := make(map[string][]string)
	for _, id := range py.pinyinTree.PrefixPredict([]byte(pinyin), 0) {
		key, err := py.pinyinTree.Key(id)
		if err != nil {
			return nil, err
		}
		value, err := py.pinyinTree.Value(id)
		if err != nil {
			return nil, err
		}
		ret[string(key)] = py.pinyinValues[value]
	}
	return ret, nil
}

func (py *Pinyin) Add(id int, keyword, pinyin, pinyinInitials string) {
	pv := make([]string, 0)
	pv = append(pv, keyword)
	py.pinyinValues = append(py.pinyinInitialsValues, pv)
	piv := make([]string, 0)
	piv = append(piv, keyword)
	py.pinyinInitialsValues = append(py.pinyinInitialsValues, piv)
	py.pinyinTree.Insert([]byte(pinyin), id)
	py.pinyinInitialsTree.Insert([]byte(pinyinInitials), id)
}

func (py *Pinyin) LoadData() error {
	curd := admin.NewPinyin()
	list, total := curd.PinyinGetList(nil, nil)
	fmt.Println(total)
	for _, item := range list {
		fmt.Println(item)
		py.pinyinInitialsTree.Insert([]byte(item.PinyinInitials), int(item.ID))
	}
	return nil
}
