package admin

import (
	"comma/pkg/library/db"

	"github.com/caicaispace/gohelper/business"
	"github.com/caicaispace/gohelper/datetime"
	"github.com/caicaispace/gohelper/syntax"
)

type HighFrequency struct {
	ID         int `gorm:"id" json:"id"`
	WordId     int `gorm:"word_id" json:"word_id"`
	ProjectId  int `gorm:"project_id" json:"project_id"`
	IsDel      int `gorm:"is_del" json:"is_del"`
	CreateTime int `gorm:"create_time" json:"create_time"`
	UpdateTime int `gorm:"update_time" json:"update_time"`
}

type HighFrequencyList struct {
	HighFrequency
	Word        string `gorm:"word" json:"word"`
	ProjectName string `gorm:"project_name" json:"project_name"`
}

const highFrequencyTableName = "dict_high_frequency"

type highFrequencyService struct{}

func NewHighFrequency() *highFrequencyService {
	return &highFrequencyService{}
}

func (hfs *highFrequencyService) HighFrequencyGetList(pager *business.Pager) ([]HighFrequencyList, int64) {
	results := make([]HighFrequencyList, 0)
	table := db.DB().Table(highFrequencyTableName)
	total := int64(0)
	table.Select("count(*)").Count(&total)
	table.Select(`
dict_high_frequency.id AS id,
dict_high_frequency.word_id AS word_id,
dict_word.word AS word,
dict_high_frequency.project_id AS project_id,
dict_project.name AS project_name,
dict_high_frequency.create_time as create_time,
dict_high_frequency.update_time AS update_time
`)
	table.Joins("left join dict_word ON dict_high_frequency.word_id = dict_word.id")
	table.Joins("left join dict_project ON dict_high_frequency.project_id = dict_project.id")
	table.Order("dict_high_frequency.id DESC")
	table.Limit(pager.GetLimit()).Offset(pager.GetOffset())
	table.Find(&results)
	return results, total
}

type HighFrequencyCreateForm struct {
	WordId    int `json:"word_id"`
	ProjectId int `json:"project_id"`
}

func (hfs *highFrequencyService) HighFrequencyCreate(inData HighFrequencyCreateForm) (*HighFrequency, error) {
	outData := HighFrequency{
		WordId:     inData.WordId,
		ProjectId:  inData.ProjectId,
		CreateTime: int(datetime.NowTimestamp()),
	}
	ret := db.DB().Table(highFrequencyTableName).Create(&outData)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return &outData, nil
}

type HighFrequencyUpdateForm struct {
	WordId     int `json:"word_id"`
	ProjectId  int `json:"project_id"`
	IsDel      int `json:"is_del"`
	UpdateTime int `json:"update_time"`
}

func (hfs *highFrequencyService) HighFrequencyUpdateById(id int, inData HighFrequencyUpdateForm) bool {
	updateData, _ := syntax.StructToMap(inData, "json")
	updateData["update_time"] = int(datetime.NowTimestamp())
	ret := db.DB().Table(highFrequencyTableName).Where("id = ?", id).Updates(updateData)
	if ret.Error != nil {
		return false
	}
	return ret.RowsAffected > 0
}

type HighFrequencyMultipleDeleteForm struct {
	Ids []int `json:"ids"`
}

func (hfs *highFrequencyService) HighFrequencyDeleteByIds(inData HighFrequencyMultipleDeleteForm) bool {
	ret := db.DB().Table(highFrequencyTableName).Where("id", inData.Ids).Update("is_del", 1)
	if ret.Error != nil {
		return false
	}
	return ret.RowsAffected > 0
}
