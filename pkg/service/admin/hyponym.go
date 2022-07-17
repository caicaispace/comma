package admin

import (
	"comma/pkg/library/db"
	"comma/pkg/library/util"
	"comma/pkg/library/util/business"
)

type Hyponym struct {
	ID             uint    `gorm:"id" json:"id"`
	HypernymWordId uint    `gorm:"hypernym_word_id" json:"hypernym_word_id"`
	HyponymWordId  uint    `gorm:"hyponym_word_id" json:"hyponym_word_id"`
	Rate           float32 `gorm:"rate" json:"rate"`
	ProjectId      uint    `gorm:"project_id" json:"project_id"`
	IsDel          int     `gorm:"is_del" json:"is_del"`
	CreateTime     int     `gorm:"create_time" json:"create_time"`
	UpdateTime     int     `gorm:"update_time" json:"update_time"`
}

type HyponymList struct {
	Hyponym
	HypernymWord string `gorm:"hypernym_word" json:"hypernym_word"`
	HyponymWord  string `gorm:"hyponym_word" json:"hyponym_word"`
	ProjectName  string `gorm:"project_name" json:"project_name"`
}

const hyponymTableName = "dict_hyponym"

type hyponymService struct{}

func NewHyponym() *hyponymService {
	return &hyponymService{}
}

func (hs *hyponymService) HyponymGetList(pager *business.Pager) ([]HyponymList, int64) {
	list := make([]HyponymList, 0)
	table := db.DB().Table(hyponymTableName)
	total := int64(0)
	table.Select("count(*)").Count(&total)
	pager.SetTotal(int(total))
	table.Select(`
dict_hyponym.id AS id,
dict_hyponym.hypernym_word_id AS word_id,
dict_hyponym.hypernym_word_id AS hypernym_word_id,
hypernym_word.word AS hypernym_word,
dict_hyponym.hyponym_word_id AS hyponym_word_id,
hyponym_word.word AS hyponym_word,
dict_hyponym.project_id AS project_id,
dict_project.name AS project_name,
dict_hyponym.create_time as create_time,
dict_hyponym.update_time AS update_time
`)
	table.Joins("left join dict_word as hypernym_word ON dict_hyponym.hypernym_word_id = hypernym_word.id")
	table.Joins("left join dict_word as hyponym_word ON dict_hyponym.hyponym_word_id = hyponym_word.id")
	table.Joins("left join dict_project ON dict_hyponym.project_id = dict_project.id")
	table.Where("dict_hyponym.is_del", 0)
	table.Order("dict_hyponym.hypernym_word_id,dict_hyponym.id DESC")
	table.Limit(pager.GetLimit()).Offset(pager.GetOffset())
	table.Find(&list)
	return list, total
}

type HyponymCreateForm struct {
	HypernymWordId uint    `json:"hypernym_word_id"`
	HyponymWordId  uint    `json:"hyponym_word_id"`
	ProjectId      uint    `json:"project_id"`
	Rate           float32 `json:"rate"`
}

func (hs *hyponymService) HyponymCreate(inData HyponymCreateForm) (*Hyponym, error) {
	outData := Hyponym{
		HypernymWordId: inData.HypernymWordId,
		ProjectId:      inData.ProjectId,
		Rate:           inData.Rate,
		HyponymWordId:  inData.HyponymWordId,
		CreateTime:     int(util.NowTimestamp()),
	}
	ret := db.DB().Table(hyponymTableName).Create(&outData)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return &outData, nil
}

type HyponymUpdateForm struct {
	HypernymWordId uint    `json:"hypernym_word_id"`
	HyponymWordId  uint    `json:"hyponym_word_id"`
	ProjectId      uint    `json:"project_id"`
	Rate           float32 `json:"rate"`
	IsDel          int     `json:"is_del"`
	UpdateTime     int     `json:"update_time"`
}

func (hs *hyponymService) HyponymUpdateById(id int, inData HyponymUpdateForm) bool {
	updateData, _ := util.StructToMap(inData, "json")
	updateData["update_time"] = int(util.NowTimestamp())
	ret := db.DB().Table(hyponymTableName).Where("id = ?", id).Updates(updateData)
	if ret.Error != nil {
		return false
	}
	return ret.RowsAffected > 0
}

type HyponymMultipleDeleteForm struct {
	Ids []int `json:"ids"`
}

func (hs *hyponymService) HyponymDeleteByIds(inData HyponymMultipleDeleteForm) bool {
	ret := db.DB().Table(hyponymTableName).Where("id", inData.Ids).Update("is_del", 1)
	if ret.Error != nil {
		return false
	}
	return ret.RowsAffected > 0
}
