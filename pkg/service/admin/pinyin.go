package admin

import (
	"github.com/caicaispace/gohelper/business"
	"github.com/caicaispace/gohelper/datetime"
	"github.com/caicaispace/gohelper/orm/gorm"
	"github.com/caicaispace/gohelper/syntax"
)

type Pinyin struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Keyword        string `gorm:"keyword" json:"keyword"`
	Pinyin         string `gorm:"pinyin" json:"pinyin"`
	Pinyins        string `gorm:"pinyins" json:"pinyins"`
	PinyinInitials string `gorm:"pinyin_initials" json:"pinyin_initials"`
	IsDel          int    `gorm:"is_del" json:"is_del"`
	CreateTime     int    `gorm:"create_time" json:"create_time"`
	UpdateTime     int    `gorm:"update_time" json:"update_time"`
}

type PinyinList struct {
	Pinyin
}

type pinyinService struct{}

const pinyinTableName = "dict_pinyin"

func NewPinyin() *pinyinService {
	return &pinyinService{}
}

func (ws *pinyinService) PinyinGetList(pager *business.Pager, filter *Pinyin) ([]*PinyinList, int64) {
	outData := make([]*PinyinList, 0)
	total := int64(0)
	table := gorm.GetInstance().GetDB("").Table(pinyinTableName)
	if (Pinyin{} != *filter) {
		if filter.Keyword != "" {
			table.Where("keyword LIKE ?", filter.Keyword+"%")
		}
		if filter.IsDel > 0 {
			table.Where("is_del", filter.IsDel)
		}
	}
	table.Select("count(*)").Count(&total)
	pager.SetTotal(int(total))
	table.Select("*")
	table.Order("id DESC")
	if (business.Pager{} != *pager) {
		table.Limit(pager.GetLimit()).Offset(pager.GetOffset())
	}
	table.Find(&outData)
	return outData, total
}

func (ws *pinyinService) PinyinGetListByIds(ids []int) ([]*PinyinList, error) {
	outData := make([]*PinyinList, 0)
	model := gorm.GetInstance().GetDB("").Table(pinyinTableName)
	ret := model.Where("id", ids).Find(&outData)
	if ret.Error != nil {
		return outData, ret.Error
	}
	return outData, nil
}

func (ws *pinyinService) PinyinGetInfoById(id int) (Pinyin, error) {
	outData := Pinyin{}
	model := gorm.GetInstance().GetDB("").Table(pinyinTableName)
	ret := model.Where("id", id).First(&outData)
	if ret.Error != nil {
		return outData, ret.Error
	}
	return outData, nil
}

func (ws *pinyinService) PinyinGetInfoByIds(ids []int) (Pinyin, error) {
	outData := Pinyin{}
	model := gorm.GetInstance().GetDB("").Table(pinyinTableName)
	ret := model.Find(&outData, ids)
	if ret.Error != nil {
		return outData, ret.Error
	}
	return outData, nil
}

type PinyinCreateForm struct {
	Keyword        string `form:"keyword" json:"keyword" valid:"MaxSize(50)"`
	Pinyin         string `gorm:"pinyin" json:"pinyin"`
	Pinyins        string `gorm:"pinyins" json:"pinyins"`
	PinyinInitials string `gorm:"pinyin_initials" json:"pinyin_initials"`
	IsDel          int    `form:"is_del" json:"is_del"`
}

func (ws *pinyinService) PinyinCreate(inData PinyinCreateForm) (*Pinyin, error) {
	// ws.Aop.RunCreateBefore(inData)
	outData := Pinyin{
		Keyword:        inData.Keyword,
		Pinyin:         inData.Pinyin,
		Pinyins:        inData.Pinyins,
		PinyinInitials: inData.PinyinInitials,
		IsDel:          inData.IsDel,
		CreateTime:     int(datetime.NowTimestamp()),
	}
	ret := gorm.GetInstance().GetDB("").Table(pinyinTableName).Create(&outData)
	if ret.Error != nil {
		return &outData, ret.Error
	}
	// ws.Aop.RunCreateAfter(outData)
	return &outData, nil
}

type PinyinUpdateForm struct {
	Keyword        string `form:"keyword" json:"keyword" valid:"Required;MaxSize(50)"`
	Pinyin         string `gorm:"pinyin" json:"pinyin"`
	Pinyins        string `gorm:"pinyins" json:"pinyins"`
	PinyinInitials string `gorm:"pinyin_initials" json:"pinyin_initials"`
	IsDel          int    `form:"is_del" json:"is_del"`
	UpdateTime     int    `json:"update_time"`
}

func (ws *pinyinService) PinyinUpdateById(id int, inData PinyinUpdateForm) bool {
	updateData, _ := syntax.StructToMap(inData, "form")
	updateData["update_time"] = int(datetime.NowTimestamp())
	ret := gorm.GetInstance().GetDB("").Table(pinyinTableName).Where("id = ?", id).Updates(updateData)
	if ret.Error != nil {
		return false
	}
	return ret.RowsAffected > 0
}

type PinyinMultipleDeleteForm struct {
	Ids []int `form:"ids" json:"ids"`
}

func (ws *pinyinService) PinyinDeleteByIds(inData PinyinMultipleDeleteForm) bool {
	ret := gorm.GetInstance().GetDB("").Table(pinyinTableName).Where("id", inData.Ids).Update("is_del", 1)
	if ret.Error != nil {
		return false
	}
	return ret.RowsAffected > 0
}
