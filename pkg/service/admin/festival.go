package admin

import (
	"errors"

	"github.com/caicaispace/gohelper/business"
	"github.com/caicaispace/gohelper/datetime"
	"github.com/caicaispace/gohelper/orm/gorm"
	"github.com/caicaispace/gohelper/syntax"
)

type Festival struct {
	ID         int    `gorm:"id" json:"id"`
	WordId     int    `gorm:"word_id" json:"word_id"`
	Name       string `gorm:"name" json:"name"`
	SunDate    string `gorm:"sun_date" json:"sun_date"`
	LunarDate  string `gorm:"lunar_date" json:"lunar_date"`
	ProjectId  int    `gorm:"project_id" json:"project_id"`
	IsDel      int    `gorm:"is_del" json:"is_del"`
	CreateTime int    `gorm:"create_time" json:"create_time"`
	UpdateTime int    `gorm:"update_time" json:"update_time"`
}

type FestivalList struct {
	Festival
	Word        string `gorm:"word" json:"word"`
	ProjectName string `gorm:"project_name" json:"project_name"`
}

const festivalTableName = "dict_festival"

type festivalService struct{}

func NewFestival() *festivalService {
	return &festivalService{}
}

func (fs *festivalService) FestivalGetList(pager *business.Pager) ([]FestivalList, int64) {
	list := make([]FestivalList, 0)
	table := gorm.GetInstance().GetDB("").Table(festivalTableName)
	total := int64(0)
	table.Select("count(*)").Count(&total)
	pager.SetTotal(int(total))
	table.Select(`
dict_festival.id AS id,
dict_festival.word_id AS word_id,
dict_word.word AS word,
dict_festival.sun_date AS sun_date,
dict_festival.lunar_date AS lunar_date,
dict_festival.project_id AS project_id,
dict_project.name AS project_name,
dict_festival.create_time as create_time,
dict_festival.update_time AS update_time
`)
	table.Joins("left join dict_word ON dict_festival.word_id = dict_word.id")
	table.Joins("left join dict_project ON dict_festival.project_id = dict_project.id")
	table.Order("dict_festival.id DESC")
	table.Limit(pager.GetLimit()).Offset(pager.GetOffset())
	table.Find(&list)
	return list, total
}

type FestivalCreateForm struct {
	WordId    int    `json:"word_id"`
	ProjectId int    `json:"project_id"`
	Name      string `json:"name"`
	SunDate   string `json:"sun_date"`
	LunarDate string `json:"lunar_date"`
	IsDel     int    `json:"is_del"`
}

func (fs *festivalService) FestivalCreate(inData FestivalCreateForm) (*Festival, error) {
	outData := Festival{
		WordId:     inData.WordId,
		ProjectId:  inData.ProjectId,
		Name:       inData.Name,
		SunDate:    inData.SunDate,
		IsDel:      inData.IsDel,
		CreateTime: int(datetime.NowTimestamp()),
	}
	ret := gorm.GetInstance().GetDB("").Table(festivalTableName).Create(&outData)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return &outData, nil
}

type FestivalUpdateForm struct {
	WordId    int    `json:"word_id"`
	ProjectId int    `json:"project_id"`
	Name      string `json:"name"`
	SunDate   string `json:"sun_date"`
	LunarDate string `json:"lunar_date"`
	IsDel     int    `json:"is_del"`
}

func (fs *festivalService) FestivalUpdateById(id int, inData FestivalUpdateForm) (*FestivalUpdateForm, error) {
	updateData, _ := syntax.StructToMap(inData, "json")
	updateData["update_time"] = int(datetime.NowTimestamp())
	ret := gorm.GetInstance().GetDB("").Table(festivalTableName).Where("id = ?", id).Updates(updateData)
	if ret.Error != nil {
		return nil, ret.Error
	}
	if ret.RowsAffected <= 0 {
		return nil, errors.New("update err")
	}
	err := syntax.MapToStruct(updateData, &inData)
	if err != nil {
		return nil, err
	}
	return &inData, nil
}

type FestivalMultipleDeleteForm struct {
	Ids []int `json:"ids"`
}

func (fs *festivalService) FestivalDeleteByIds(inData FestivalMultipleDeleteForm) error {
	ret := gorm.GetInstance().GetDB("").Table(festivalTableName).Where("id", inData.Ids).Update("is_del", 1)
	if ret.Error != nil {
		return ret.Error
	}
	if ret.RowsAffected <= 0 {
		return errors.New("update error")
	}
	return nil
}
