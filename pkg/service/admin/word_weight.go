package admin

import (
	"comma/pkg/library/db"
	"comma/pkg/library/util"

	"github.com/caicaispace/gohelper/business"
)

type WordWeight struct {
	ID         int     `gorm:"id" json:"id"`
	WordId     int     `gorm:"word_id" json:"word_id"`
	Weight     float32 `gorm:"weight" json:"weight"`
	ProjectId  int     `gorm:"project_id" json:"project_id"`
	IsDel      int     `gorm:"is_del" json:"is_del"`
	CreateTime int     `gorm:"create_time" json:"create_time"`
	UpdateTime int     `gorm:"update_time" json:"update_time"`
}

type WordWeightList struct {
	WordWeight
	Word string `gorm:"word" json:"word"`
}

const weightTableName = "dict_weight"

type wordWeightService struct{}

func NewWordWeight() *wordWeightService {
	return &wordWeightService{}
}

func (wws *wordWeightService) WordWeightGetList(pager *business.Pager, filter *Word) ([]WordWeightList, int64) {
	list := make([]WordWeightList, 0)
	total := int64(0)
	table := db.DB().Table(weightTableName)
	table.Select("count(*)").Count(&total)
	pager.SetTotal(int(total))
	table.Select(`
dict_weight.id AS id,
dict_weight.word_id AS word_id,
dict_word.word AS word,
dict_weight.project_id AS project_id,
dict_project.name AS project_name,
dict_weight.weight AS weight,
dict_weight.is_del AS is_del,
dict_weight.create_time as create_time,
dict_weight.update_time AS update_time
`)
	table.Joins("left join dict_word ON dict_weight.word_id = dict_word.id")
	table.Joins("left join dict_project ON dict_weight.project_id = dict_project.id")
	if (Word{} != *filter) {
		if filter.Word != "" {
			table.Where("dict_word.word LIKE ?", filter.Word+"%")
		}
		if filter.IsDel > 0 {
			table.Where("dict_weight.is_del", filter.IsDel)
		}
	}
	table.Order("dict_weight.id DESC")
	table.Limit(pager.GetLimit()).Offset(pager.GetOffset())
	table.Find(&list)
	return list, total
}

type WordWeightCreateForm struct {
	WordId    int     `json:"word_id"`
	Weight    float32 `json:"weight"`
	ProjectId int     `json:"project_id"`
}

func (wws *wordWeightService) WordWeightCreate(inData WordWeightCreateForm) (*WordWeight, error) {
	outData := WordWeight{
		WordId:     inData.WordId,
		Weight:     inData.Weight,
		ProjectId:  inData.ProjectId,
		IsDel:      1,
		CreateTime: int(util.NowTimestamp()),
	}
	ret := db.DB().Table(weightTableName).Create(&outData)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return &outData, nil
}

type WordWeightUpdateForm struct {
	WordId     int     `json:"word_id"`
	Weight     float32 `json:"weight"`
	ProjectId  int     `json:"project_id"`
	IsDel      int     `json:"is_del"`
	UpdateTime int     `json:"update_time"`
}

func (wws *wordWeightService) WordWeightUpdateById(id int, inData WordWeightUpdateForm) error {
	updateData, _ := util.StructToMap(inData, "json")
	updateData["update_time"] = int(util.NowTimestamp())
	ret := db.DB().Table(weightTableName).Where("id = ?", id).Updates(updateData)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}

type WordWeightMultipleDeleteForm struct {
	Ids []int `json:"ids"`
}

func (wws *wordWeightService) WordWeightDeleteByIds(inData WordWeightMultipleDeleteForm) error {
	ret := db.DB().Table(weightTableName).Where("id", inData.Ids).Update("is_del", 1)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}
