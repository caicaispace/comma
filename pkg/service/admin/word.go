package admin

import (
	"goaway/pkg/library/db"
	"goaway/pkg/library/util"
	"goaway/pkg/library/util/business"
)

type Word struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Word       string `gorm:"word" json:"word"`
	Frequency  int    `gorm:"frequency" json:"frequency"`
	Classify   string `gorm:"classify" json:"classify"`
	IsDel      int    `gorm:"is_del" json:"is_del"`
	CreateTime int    `gorm:"create_time" json:"create_time"`
	UpdateTime int    `gorm:"update_time" json:"update_time"`
}

//func (w *Word) AfterFind(tx *gorm.DB) (err error) {
//	fmt.Println(w.ID)
//	return
//}

// func (w *Word) AfterCreate(tx *gorm.DB) (err error) {
// 	fmt.Println(w.ID)
// 	return
// }

// func (w *Word) AfterUpdate(tx *gorm.DB) (err error) {
// 	fmt.Println(w.Word)
// 	return
// }

type WordList struct {
	Word
}

type wordService struct {
	// Aop *util.Aop
}

const wordTableName = "dict_word"

// const wordTableName = "dict_word_copy"

func NewWord() *wordService {
	return &wordService{
		//Aop: &util.Aop{
		//	AopMethod: make(map[string]func(interface{})),
		//},
	}
}

func (ws *wordService) WordGetList(pager *business.Pager, filter *Word) ([]*WordList, int64) {
	outData := make([]*WordList, 0)
	total := int64(0)
	table := db.DB().Table(wordTableName)
	if (Word{} != *filter) {
		if filter.Word != "" {
			table.Where("word LIKE ?", filter.Word+"%")
		}
		if filter.Frequency > 0 {
			table.Where("frequency", filter.Frequency)
		}
		if filter.Classify != "" {
			table.Where("classify", filter.Classify)
		}
		if filter.IsDel > 0 {
			table.Where("is_del", filter.IsDel)
		}
	}
	table.Select("count(*)").Count(&total)
	pager.SetTotal(int(total))
	table.Select("*")
	table.Order("id DESC")
	table.Limit(pager.GetLimit()).Offset(pager.GetOffset())
	table.Find(&outData)
	return outData, total
}

func (ws *wordService) WordGetListByIds(ids []int) ([]*WordList, error) {
	outData := make([]*WordList, 0)
	model := db.DB().Table(wordTableName)
	ret := model.Where("id", ids).Find(&outData)
	if ret.Error != nil {
		return outData, ret.Error
	}
	return outData, nil
}

func (ws *wordService) WordGetInfoById(id int) (Word, error) {
	outData := Word{}
	model := db.DB().Table(wordTableName)
	ret := model.Where("id", id).First(&outData)
	if ret.Error != nil {
		return outData, ret.Error
	}
	return outData, nil
}

func (ws *wordService) WordGetInfoByIds(ids []int) (Word, error) {
	outData := Word{}
	model := db.DB().Table(wordTableName)
	ret := model.Find(&outData, ids)
	if ret.Error != nil {
		return outData, ret.Error
	}
	return outData, nil
}

type WordCreateForm struct {
	Word      string `form:"word" json:"word" valid:"MaxSize(50)"`
	Frequency int    `form:"frequency" json:"frequency"`
	Classify  string `form:"classify" json:"classify"`
	IsDel     int    `json:"is_del"`
}

func (ws *wordService) WordCreate(inData WordCreateForm) (*Word, error) {
	// ws.Aop.RunCreateBefore(inData)
	outData := Word{
		Word:       inData.Word,
		Frequency:  inData.Frequency,
		Classify:   inData.Classify,
		IsDel:      inData.IsDel,
		CreateTime: int(util.NowTimestamp()),
	}
	ret := db.DB().Table(wordTableName).Create(&outData)
	if ret.Error != nil {
		return &outData, ret.Error
	}
	// ws.Aop.RunCreateAfter(outData)
	return &outData, nil
}

type WordUpdateForm struct {
	Word       string `form:"word" json:"word" valid:"Required;MaxSize(50)"`
	Frequency  int    `form:"frequency" json:"frequency"`
	Classify   string `form:"classify" json:"classify"`
	IsDel      int    `form:"is_del" json:"is_del"`
	UpdateTime int    `json:"update_time"`
}

func (ws *wordService) WordUpdateById(id int, inData WordUpdateForm) bool {
	updateData, _ := util.StructToMap(inData, "form")
	updateData["update_time"] = int(util.NowTimestamp())
	ret := db.DB().Table(wordTableName).Where("id = ?", id).Updates(updateData)
	if ret.Error != nil {
		return false
	}
	return ret.RowsAffected > 0
}

type WordMultipleDeleteForm struct {
	Ids []int `form:"ids" json:"ids"`
}

func (ws *wordService) WordDeleteByIds(inData WordMultipleDeleteForm) bool {
	ret := db.DB().Table(wordTableName).Where("id", inData.Ids).Update("is_del", 1)
	if ret.Error != nil {
		return false
	}
	return ret.RowsAffected > 0
}
