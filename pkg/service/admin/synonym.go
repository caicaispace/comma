package admin

import (
	"errors"

	"github.com/caicaispace/gohelper/business"
	"github.com/caicaispace/gohelper/datetime"
	"github.com/caicaispace/gohelper/orm/gorm"
	"github.com/caicaispace/gohelper/syntax"
)

type Synonym struct {
	ID         int     `gorm:"id" json:"id"`
	WordIds    string  `gorm:"word_ids" json:"word_ids"`
	ProjectId  int     `gorm:"project_id" json:"project_id"`
	Rate       float32 `gorm:"rate" json:"rate"`
	IsDel      int     `gorm:"is_del" json:"is_del"`
	CreateTime int     `gorm:"create_time" json:"create_time"`
	UpdateTime int     `gorm:"update_time" json:"update_time"`
}

type SynonymList struct {
	Synonym
	ProjectName string `gorm:"project_name" json:"project_name"`
}

const synonymTableName = "dict_synonyms"

type synonymService struct{}

func NewSynonym() *synonymService {
	return &synonymService{}
}

func (ss *synonymService) SynonymGetList(pager *business.Pager) ([]SynonymList, int64) {
	list := make([]SynonymList, 0)
	table := gorm.GetInstance().GetDB("").Table(synonymTableName)
	total := int64(0)
	table.Select("count(*)").Count(&total)
	pager.SetTotal(int(total))
	table.Select(`
dict_synonyms.id AS id,
dict_synonyms.word_ids AS word_ids,
dict_synonyms.project_id AS project_id,
dict_project.name AS project_name,
dict_synonyms.rate AS rate,
dict_synonyms.create_time as create_time,
dict_synonyms.update_time AS update_time
`)
	table.Joins("left join dict_project ON dict_synonyms.project_id = dict_project.id")
	table.Where("dict_synonyms.is_del", 0)
	table.Order("dict_synonyms.id DESC")
	table.Limit(pager.GetLimit()).Offset(pager.GetOffset())
	table.Find(&list)
	return list, total
}

type SynonymCreateForm struct {
	WordIds   string  `json:"word_ids"`
	ProjectId int     `json:"project_id"`
	IsDel     int     `json:"is_del"`
	Rate      float32 `json:"rate"`
}

func (ss *synonymService) SynonymCreate(inData SynonymCreateForm) (*Synonym, error) {
	outData := Synonym{
		WordIds:    inData.WordIds,
		Rate:       inData.Rate,
		ProjectId:  inData.ProjectId,
		IsDel:      inData.IsDel,
		CreateTime: int(datetime.NowTimestamp()),
	}
	ret := gorm.GetInstance().GetDB("").Table(synonymTableName).Create(&outData)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return &outData, nil
}

type SynonymUpdateForm struct {
	WordIds    string  `json:"word_ids"`
	ProjectId  int     `json:"project_id"`
	Rate       float32 `json:"rate"`
	IsDel      int     `json:"is_del"`
	UpdateTime int     `json:"update_time"`
}

func (ss *synonymService) SynonymUpdateById(id int, inData SynonymUpdateForm) (*SynonymUpdateForm, error) {
	updateData, _ := syntax.StructToMap(inData, "json")
	updateData["update_time"] = int(datetime.NowTimestamp())
	ret := gorm.GetInstance().GetDB("").Table(synonymTableName).Where("id = ?", id).Updates(updateData)
	if ret.RowsAffected <= 0 {
		return nil, errors.New("update err")
	}
	err := syntax.MapToStruct(updateData, &inData)
	if err != nil {
		return nil, err
	}
	return &inData, nil
}

type SynonymUMultipleDeleteForm struct {
	Ids []int `json:"ids"`
}

func (ss *synonymService) SynonymDeleteByIds(inData SynonymUMultipleDeleteForm) bool {
	ret := gorm.GetInstance().GetDB("").Table(synonymTableName).Where("id", inData.Ids).Update("is_del", 1)
	if ret.Error != nil {
		return false
	}
	return ret.RowsAffected > 0
}
