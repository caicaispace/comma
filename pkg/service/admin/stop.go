package admin

import (
	"comma/pkg/library/db"
	"errors"

	"github.com/caicaispace/gohelper/business"
	"github.com/caicaispace/gohelper/datetime"
	"github.com/caicaispace/gohelper/syntax"
)

type Stop struct {
	ID         int `gorm:"id" json:"id"`
	WordId     int `gorm:"word_id" json:"word_id"`
	ProjectId  int `gorm:"project_id" json:"project_id"`
	IsDel      int `gorm:"is_del" json:"is_del"`
	CreateTime int `gorm:"create_time" json:"create_time"`
	UpdateTime int `gorm:"update_time" json:"update_time"`
}

type StopList struct {
	Stop
	Word        string `gorm:"word" json:"word"`
	ProjectName string `gorm:"project_name" json:"project_name"`
}

const stopTableName = "dict_stop"

type stopService struct{}

func NewStop() *stopService {
	return &stopService{}
}

func (ss *stopService) StopGetList(pager *business.Pager, filter *Word) ([]StopList, int64) {
	list := make([]StopList, 0)
	table := db.DB().Table(stopTableName)
	total := int64(0)
	table.Select("count(*)").Count(&total)
	pager.SetTotal(int(total))
	table.Select(`
dict_stop.id AS id,
dict_stop.word_id AS word_id,
dict_word.word AS word,
dict_stop.project_id AS project_id,
dict_project.name AS project_name,
dict_stop.is_del AS is_del,
dict_stop.create_time as create_time,
dict_stop.update_time AS update_time
`)
	table.Joins("left join dict_word ON dict_stop.word_id = dict_word.id")
	table.Joins("left join dict_project ON dict_stop.project_id = dict_project.id")
	if (Word{} != *filter) {
		if filter.Word != "" {
			table.Where("dict_word.word LIKE ?", filter.Word+"%")
		}
		if filter.IsDel > 0 {
			table.Where("dict_banned.is_del", filter.IsDel)
		}
	}
	table.Order("dict_stop.id DESC")
	table.Limit(pager.GetLimit()).Offset(pager.GetOffset())
	table.Find(&list)
	return list, total
}

type StopCreateForm struct {
	WordId    int `json:"word_id"`
	ProjectId int `json:"project_id"`
}

func (ss *stopService) StopCreate(inData StopCreateForm) (*Stop, error) {
	outData := Stop{
		WordId:     inData.WordId,
		ProjectId:  inData.ProjectId,
		CreateTime: int(datetime.NowTimestamp()),
	}
	//if inData.ProjectId > 0 {
	//	projectService := NewProject()
	//	projectModel, err := projectService.ProjectGetInfoById(inData.ProjectId)
	//	if err != nil {
	//		return outData, err
	//	}
	//	outData.Name = projectModel.Name
	//}
	ret := db.DB().Table(stopTableName).Create(&outData)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return &outData, nil
}

type StopUpdateForm struct {
	WordId     int `json:"word_id"`
	ProjectId  int `json:"project_id"`
	IsDel      int `json:"is_del"`
	UpdateTime int `json:"update_time"`
}

func (ss *stopService) StopUpdateById(id int, inData StopUpdateForm) (*StopUpdateForm, error) {
	updateData, _ := syntax.StructToMap(inData, "json")
	updateData["update_time"] = int(datetime.NowTimestamp())
	ret := db.DB().Table(stopTableName).Where("id = ?", id).Updates(updateData)
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

type StopMultipleDeleteForm struct {
	Ids []int `json:"ids" json:"ids"`
}

func (ss *stopService) StopDeleteByIds(inData StopMultipleDeleteForm) bool {
	ret := db.DB().Table(stopTableName).Where("id", inData.Ids).Update("is_del", 1)
	if ret.Error != nil {
		return false
	}
	return ret.RowsAffected > 0
}
