package banned

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/caicaispace/gohelper/logx"
	"github.com/caicaispace/gohelper/orm/gorm"
	orm "gorm.io/gorm"
)

var f = NewFilter()

type Banned struct{}

var (
	service *Banned
	once    sync.Once
)

func GetInstance() *Banned {
	once.Do(func() {
		service = &Banned{}
	})
	return service
}

func New() *Banned {
	return &Banned{}
}

func (fs *Banned) LoadData() {
	var err error
	var words []string
	defer func() {
		// words variable gc
		words = nil
	}()
	words, err = _loadDataFromDb(10)
	if err != nil {
		logx.Error(err)
		os.Exit(0)
	}
	f.AddWord(words...)
}

func (fs *Banned) Find(word, handleType string) (bool, []string) {
	textFindSlice := []string{""}
	hasFind := false
	switch handleType {
	case "replace":
		textFindSlice = []string{f.Replace(word, '*')}
		hasFind = len(textFindSlice) > 0
	case "findIn":
		has, text := f.FindIn(word)
		textFindSlice = []string{text}
		hasFind = has
	case "findAll":
		textFindSlice = f.FindAll(word)
		hasFind = len(textFindSlice) > 0
	}
	return hasFind, textFindSlice
}

func (fs *Banned) Add(word string) {
	f.AddWord(word)
}

func (fs *Banned) Del(word string) {
	f.DelWord(word)
}

func _loadDataFromDb(limit int) ([]string, error) {
	type Model struct {
		ID   uint   `gorm:"primaryKey"`
		Word string `gorm:"Word"`
		Must uint   `gorm:"must"`
	}
	outData := make([]string, 0)
	model := gorm.GetInstance().GetDB("").Table("banned")
	model.Select("*")
	model.Limit(limit)
	model.Order("id DESC")
	rows := make([]*Model, 0)
	model.FindInBatches(&rows, 5000, func(tx *orm.DB, batch int) error {
		for _, row := range rows {
			outData = append(outData, row.Word)
		}
		return nil
	})
	return outData, nil
}

func _loadDataFromFile() ([]string, error) {
	outData := make([]string, 0)
	// err = f.LoadWordDict(setting.FilterSetting.LoadDictPath)
	err := f.LoadNetWordDict("https://raw.githubusercontent.com/importcjj/sensitive/master/dict/dict.txt?t=" + strconv.Itoa(int(time.Now().Unix())))
	if err != nil {
		return outData, err
	}
	return outData, nil
}
